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

package operationButton

import (
	"context"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/modules/cmp"
)

type ComponentOperationButton struct {
	ctx    context.Context
	sdk    *cptype.SDK
	server cmp.SteveServer
	Type   string `json:"type,omitempty"`
	Props  Props  `json:"props"`
}

type Props struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Menu []Menu `json:"menu,omitempty"`
}
type Menu struct {
	Key        string                 `json:"key,omitempty"`
	Text       string                 `json:"text,omitempty"`
	Operations map[string]interface{} `json:"operations,omitempty"`
}

type Operation struct {
	Key        string  `json:"key,omitempty"`
	Reload     bool    `json:"reload"`
	Confirm    string  `json:"confirm,omitempty"`
	Command    Command `json:"command,omitempty"`
	SuccessMsg string  `json:"successMsg,omitempty"`
}

type Command struct {
	Key    string       `json:"key,omitempty"`
	Target string       `json:"target,omitempty"`
	State  CommandState `json:"state,omitempty"`
}

type CommandState struct {
	Params map[string]string `json:"params,omitempty"`
}
