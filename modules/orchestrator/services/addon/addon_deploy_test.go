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

package addon

import (
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/orchestrator/dbclient"
	"github.com/erda-project/erda/pkg/parser/diceyml"
)

func TestMysqlPreProcess(t *testing.T) {
	params := &apistructs.AddonHandlerCreateItem{Plan: apistructs.AddonBasic}
	addonSpec := &apistructs.AddonExtension{Plan: map[string]apistructs.AddonPlanItem{
		apistructs.AddonBasic: {
			Nodes: 2,
		},
	}}
	addonDeployGroup := &apistructs.ServiceGroupCreateV2Request{
		GroupLabels: map[string]string{
			"ADDON_GROUPS": "2",
		},
	}
	mysqlPreProcess(params, addonSpec, addonDeployGroup)
	assert.Equal(t, "1", addonDeployGroup.GroupLabels["ADDON_GROUPS"])
	assert.Equal(t, 1, addonSpec.Plan[apistructs.AddonBasic].Nodes)
}

func TestBuildAddonRequestGroupForEsWithOperator(t *testing.T) {
	defer monkey.UnpatchAll()

	bdl := bundle.New()

	monkey.PatchInstanceMethod(reflect.TypeOf(bdl), "CapacityInfo", func(bundle *bundle.Bundle, clusterName string) (*apistructs.CapacityInfoResponse, error) {
		return &apistructs.CapacityInfoResponse{
			Data: apistructs.CapacityInfoData{
				ElasticsearchOperator: true,
			},
		}, nil
	})

	monkey.PatchInstanceMethod(reflect.TypeOf(bdl), "QueryClusterInfo", func(bundle *bundle.Bundle, clusterName string) (apistructs.ClusterInfoData, error) {
		return apistructs.ClusterInfoData{}, nil
	})

	addon := New(WithBundle(bdl))

	_, err := addon.BuildAddonRequestGroup(&apistructs.AddonHandlerCreateItem{
		AddonName: apistructs.AddonES,
	}, &dbclient.AddonInstance{
		AddonID: "fake-addon-id",
	}, &apistructs.AddonExtension{
		Version: "6.8.8",
	}, &diceyml.Object{
		Services: diceyml.Services{
			"fake-service1": &diceyml.Service{},
			"fake-service2": &diceyml.Service{},
			"fake-service3": &diceyml.Service{},
		},
	})

	assert.NoError(t, err)
}
