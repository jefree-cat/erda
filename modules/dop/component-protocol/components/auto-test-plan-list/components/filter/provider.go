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

package filter

import (
	"github.com/erda-project/erda-infra/providers/component-protocol/components/filter"
	"github.com/erda-project/erda-infra/providers/component-protocol/components/filter/impl"
	model "github.com/erda-project/erda-infra/providers/component-protocol/components/filter/models"
	"github.com/erda-project/erda-infra/providers/component-protocol/cpregister"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/dop/component-protocol/components/auto-test-plan-list/i18n"
	"github.com/erda-project/erda/modules/dop/component-protocol/types"
)

func init() {
	cpregister.RegisterComponent("auto-test-plan-list", "filter", func() cptype.IComponent { return &ComponentFilter{} })
}

type ComponentFilter struct {
	impl.DefaultFilter
	sdk       *cptype.SDK
	bdl       *bundle.Bundle
	State     State
	projectID uint64
}

type State struct {
	Archive   *bool    `json:"archive"`
	Name      string   `json:"name"`
	Iteration []uint64 `json:"iteration"`
	Values    Value    `json:"values,omitempty"`
}

type Value struct {
	Archive   []string `json:"archive"`
	Name      string   `json:"name"`
	Iteration []uint64 `json:"iteration"`
}

func (f *ComponentFilter) BeforeHandleOp(sdk *cptype.SDK) {
	f.bdl = sdk.Ctx.Value(types.GlobalCtxKeyBundle).(*bundle.Bundle)
	f.sdk = sdk
	cputil.MustObjJSONTransfer(&f.StdStatePtr, &f.State)
	projectID := cputil.GetInParamByKey(sdk.Ctx, "projectId").(float64)
	f.projectID = uint64(projectID)
}

func (f *ComponentFilter) RegisterInitializeOp() (opFunc cptype.OperationFunc) {
	return func(sdk *cptype.SDK) cptype.IStdStructuredPtr {
		iterations, err := f.bdl.ListProjectIterations(apistructs.IterationPagingRequest{
			PageNo:              1,
			PageSize:            999,
			ProjectID:           f.projectID,
			WithoutIssueSummary: true,
		}, "0")
		if err != nil {
			panic(err)
		}
		var options []model.SelectOption
		for _, iteration := range iterations {
			options = append(options, *model.NewSelectOption(iteration.Title, iteration.ID))
		}
		iteration := model.NewSelectCondition("iteration", sdk.I18n(i18n.I18nKeyIteration), options)

		name := model.NewInputCondition("name", sdk.I18n(i18n.I18nKeyPlanName))
		name.Placeholder = sdk.I18n("searchPlan")
		name.Outside = true
		archive := model.NewSelectCondition("archive", sdk.I18n(i18n.I18nKeyArchive), []model.SelectOption{
			*model.NewSelectOption(sdk.I18n(i18n.I18nKeyInProgress), "inprogress"),
			*model.NewSelectOption(sdk.I18n(i18n.I18nKeyArchived), "archived"),
		})

		f.StdDataPtr.Conditions = []interface{}{
			name, archive, iteration,
		}
		f.StdDataPtr.HideSave = true
		f.StdDataPtr.Operations = map[cptype.OperationKey]cptype.Operation{
			filter.OpFilter{}.OpKey(): cputil.NewOpBuilder().Build(),
		}
		f.State.Values = Value{
			Archive: []string{"inprogress"},
		}
		archived := false
		f.State.Archive = &archived
		return nil
	}
}

func (f *ComponentFilter) RegisterRenderingOp() (opFunc cptype.OperationFunc) {
	return f.RegisterInitializeOp()
}

func (f *ComponentFilter) RegisterFilterOp(opData filter.OpFilter) (opFunc cptype.OperationFunc) {
	return func(sdk *cptype.SDK) cptype.IStdStructuredPtr {
		f.State.Name = f.State.Values.Name
		f.State.Iteration = f.State.Values.Iteration
		f.State.Archive = nil
		if len(f.State.Values.Archive) == 1 {
			archived := f.State.Values.Archive[0] == "archived"
			f.State.Archive = &archived
		}
		return nil
	}
}

func (f *ComponentFilter) AfterHandleOp(sdk *cptype.SDK) {
	cputil.MustObjJSONTransfer(&f.State, &f.StdStatePtr)
}

func (f *ComponentFilter) RegisterFilterItemSaveOp(opData filter.OpFilterItemSave) (opFunc cptype.OperationFunc) {
	return func(sdk *cptype.SDK) cptype.IStdStructuredPtr {
		return nil
	}
}

func (f *ComponentFilter) RegisterFilterItemDeleteOp(opData filter.OpFilterItemDelete) (opFunc cptype.OperationFunc) {
	return func(sdk *cptype.SDK) cptype.IStdStructuredPtr {
		return nil
	}
}
