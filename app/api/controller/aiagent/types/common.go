// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"github.com/harness/gitness/app/api/controller/aiagent/types/enum"
)

type Suggestion struct {
	ID             string
	Prompt         string
	UserSuggestion string
	Suggestion     string
}

type Conversation struct {
	Role    enum.Role `json:"role"`
	Message Message   `json:"message"`
}

type Message struct {
	Type enum.MessageType `json:"type"`
	Data string           `json:"data"`
}

// Additional common structs can be defined here as needed.
