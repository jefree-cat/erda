// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dispatcher

import (
	"context"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"
	"gopkg.in/igm/sockjs-go.v2/sockjs"

	"github.com/erda-project/erda-infra/base/version"
	"github.com/erda-project/erda-proto-go/core/messenger/notify/pb"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/core-services/services/dingtalk/api/interfaces"
	"github.com/erda-project/erda/modules/eventbox/conf"
	"github.com/erda-project/erda/modules/eventbox/input"
	httpinput "github.com/erda-project/erda/modules/eventbox/input/http"
	"github.com/erda-project/erda/modules/eventbox/monitor"
	"github.com/erda-project/erda/modules/eventbox/register"
	"github.com/erda-project/erda/modules/eventbox/server"
	stypes "github.com/erda-project/erda/modules/eventbox/server/types"
	"github.com/erda-project/erda/modules/eventbox/subscriber"
	dingdingsubscriber "github.com/erda-project/erda/modules/eventbox/subscriber/dingding"
	dingdingworknoticesubscriber "github.com/erda-project/erda/modules/eventbox/subscriber/dingding_worknotice"
	"github.com/erda-project/erda/modules/eventbox/subscriber/dingtalk_worknotice"
	emailsubscriber "github.com/erda-project/erda/modules/eventbox/subscriber/email"
	fakesubscriber "github.com/erda-project/erda/modules/eventbox/subscriber/fake"
	groupsubscriber "github.com/erda-project/erda/modules/eventbox/subscriber/group"
	httpsubscriber "github.com/erda-project/erda/modules/eventbox/subscriber/http"
	mbox "github.com/erda-project/erda/modules/eventbox/subscriber/mbox"
	smssubscriber "github.com/erda-project/erda/modules/eventbox/subscriber/sms"
	vmssubscriber "github.com/erda-project/erda/modules/eventbox/subscriber/vms"
	"github.com/erda-project/erda/modules/eventbox/webhook"
	"github.com/erda-project/erda/modules/eventbox/websocket"
	"github.com/erda-project/erda/modules/pkg/user"
	"github.com/erda-project/erda/pkg/goroutinepool"
)

type Dispatcher interface {
	RegisterSubscriber(s subscriber.Subscriber)
	RegisterInput(input.Input)
	Start() // block until stopped
	Stop()  // block until stopped
}

type DispatcherImpl struct {
	subscribers     map[string]subscriber.Subscriber
	subscriberspool map[string]*goroutinepool.GoroutinePool
	router          *Router
	register        register.Register
	inputs          []input.Input
	httpserver      *server.Server

	runningWg sync.WaitGroup
}

func New(dingtalk interfaces.DingTalkApiClientFactory, messenger pb.NotifyServiceServer) (Dispatcher, error) {
	dispatcher := DispatcherImpl{
		subscribers:     make(map[string]subscriber.Subscriber),
		subscriberspool: make(map[string]*goroutinepool.GoroutinePool),
	}

	httpi, err := httpinput.New()
	if err != nil {
		return nil, err
	}
	wsi, err := websocket.New()
	if err != nil {
		return nil, err
	}
	fakeS, err := fakesubscriber.New(fakesubscriber.FakeTestFilePath)
	if err != nil {
		return nil, err
	}
	httpS := httpsubscriber.New()
	bundleS := bundle.New(bundle.WithCoreServices())
	dingdingS := dingdingsubscriber.New(conf.Proxy(), messenger)
	dingdingWorknoticeS := dingdingworknoticesubscriber.New(conf.Proxy(), messenger)
	mboxS := mbox.New(bundle.New(bundle.WithCoreServices()), messenger)
	emailS := emailsubscriber.New(conf.SmtpHost(), conf.SmtpPort(), conf.SmtpUser(), conf.SmtpPassword(),
		conf.SmtpDisplayUser(), conf.SmtpIsSSL(), conf.SMTPInsecureSkipVerify(), bundleS, messenger)
	smsS := smssubscriber.New(
		conf.AliyunAccessKeyID(),
		conf.AliyunAccessKeySecret(),
		conf.AliyunSmsSignName(),
		conf.AliyunSmsMonitorTemplateCode(), bundleS, messenger)
	vmsS := vmssubscriber.New(conf.AliyunAccessKeyID(), conf.AliyunAccessKeySecret(), conf.AliyunVmsMonitorTtsCode(),
		conf.AliyunVmsMonitorCalledShowNumber(), bundleS, messenger)
	dingWorkNotice := dingtalk_worknotice.New(bundleS, dingtalk, messenger)
	groupS := groupsubscriber.New(bundleS)
	if err != nil {
		return nil, err
	}

	wh, err := webhook.NewWebHookHTTP()
	if err != nil {
		return nil, err
	}
	mon, err := monitor.NewMonitorHTTP()
	if err != nil {
		return nil, err
	}

	dispatcher.RegisterInput(httpi)
	dispatcher.RegisterInput(wsi)

	dispatcher.RegisterSubscriber(fakeS)
	dispatcher.RegisterSubscriber(httpS)
	dispatcher.RegisterSubscriber(dingdingS)
	dispatcher.RegisterSubscriber(dingdingWorknoticeS)
	dispatcher.RegisterSubscriber(smsS)
	dispatcher.RegisterSubscriber(emailS)
	dispatcher.RegisterSubscriber(vmsS)
	dispatcher.RegisterSubscriber(mboxS)
	dispatcher.RegisterSubscriber(groupS)
	dispatcher.RegisterSubscriber(dingWorkNotice)

	for name := range dispatcher.subscribers {
		dispatcher.subscriberspool[name] = goroutinepool.New(conf.PoolSize())
	}

	reg, err := register.New()
	if err != nil {
		return nil, err
	}
	dispatcher.register = reg

	server, err := server.New()
	if err != nil {
		return nil, err
	}
	regHTTP := register.NewHTTP(reg)
	// add endpoints here
	server.AddEndPoints(httpi.GetHTTPEndPoints())
	server.AddEndPoints(regHTTP.GetHTTPEndPoints())
	server.AddEndPoints([]stypes.Endpoint{{"/version", http.MethodGet, getVersion}})
	server.AddEndPoints([]stypes.Endpoint{{"/actions/get-smtp-info", http.MethodGet, getSMTPInfo}})
	server.AddEndPoints(wh.GetHTTPEndPoints())
	server.AddEndPoints(mon.GetHTTPEndPoints())
	// add router for Websocket
	server.Router().PathPrefix("/api/dice/eventbox").Path("/ws/{any:.*}").
		Handler(sockjs.NewHandler("/api/dice/eventbox/ws", sockjs.DefaultOptions, wsi.HTTPHandle))
	dispatcher.httpserver = server

	router, err := NewRouter(&dispatcher)
	if err != nil {
		return nil, err
	}
	groupS.SetRoute(router)
	dispatcher.router = router

	return &dispatcher, nil
}

func (d *DispatcherImpl) GetRegister() register.Register {
	return d.register
}

func (d *DispatcherImpl) GetSubscribers() map[string]subscriber.Subscriber {
	return d.subscribers
}

func (d *DispatcherImpl) GetSubscribersPool() map[string]*goroutinepool.GoroutinePool {
	return d.subscriberspool
}
func (d *DispatcherImpl) RegisterSubscriber(s subscriber.Subscriber) {
	logrus.Infof("Dispatcher: register subscriber [%s]", s.Name())
	d.subscribers[s.Name()] = s
}

func (d *DispatcherImpl) RegisterInput(i input.Input) {
	logrus.Infof("Dispatcher: register input [%s]", i.Name())
	d.inputs = append(d.inputs, i)
}

func (d *DispatcherImpl) Start() {
	var err error
	for _, pool := range d.subscriberspool {
		pool.Start()
	}
	d.runningWg.Add(len(d.inputs) + 1)
	for _, i := range d.inputs {
		go func(i input.Input) {
			err = i.Start(d.router.Route)
			if err != nil {
				logrus.Errorf("dispatcher: start %s err:%v", i.Name(), err)
			}
			d.runningWg.Done()
		}(i)
	}
	// start httpserver
	go func() {
		if err := d.httpserver.Start(); err != nil {
			logrus.Errorf("dispatcher: start httpserver: %v", err)
		}
		d.runningWg.Done()
	}()
	d.runningWg.Wait()
}

// 1. 关闭所有输入端 (httpserver, inputs)
// 2. 等待 pool 里的所有消息发送完
// 3. 关闭 pool
// 4. 关闭 register
func (d *DispatcherImpl) Stop() {
	logrus.Info("Dispatcher: stopping")
	defer logrus.Info("Dispatcher: stopped")
	// stop httpserver first
	if err := d.httpserver.Stop(); err != nil {
		logrus.Errorf("dispatcher: stop httpserver: %v", err)
	}
	// stop inputs
	for _, i := range d.inputs {
		if err := i.Stop(); err != nil {
			logrus.Errorf("dispatcher: stop %s err:%v", i.Name(), err)
		}
	}

	// it will block until all things done
	for _, pool := range d.subscriberspool {
		pool.Stop()
	}
}

func getVersion(ctx context.Context, req *http.Request, vars map[string]string) (stypes.Responser, error) {
	return stypes.HTTPResponse{Status: http.StatusOK, Content: version.String()}, nil
}

func getSMTPInfo(ctx context.Context, req *http.Request, vars map[string]string) (stypes.Responser, error) {
	identityInfo, err := user.GetIdentityInfo(req)

	if err != nil {
		logrus.Errorf("getIdentityInfo error %v", err)
		return stypes.HTTPResponse{Status: http.StatusUnauthorized, Content: "failed"}, nil
	}

	if !identityInfo.IsInternalClient() {
		return stypes.HTTPResponse{Status: http.StatusUnauthorized, Content: "failed"}, nil
	}

	return stypes.HTTPResponse{Status: http.StatusOK, Content: emailsubscriber.NewMailSubscriberInfo(conf.SmtpHost(), conf.SmtpPort(), conf.SmtpUser(), conf.SmtpPassword(),
		conf.SmtpDisplayUser(), conf.SmtpIsSSL(), conf.SMTPInsecureSkipVerify())}, nil
}
