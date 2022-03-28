// Copyright 2022 The sacloud/phy-service-go Authors
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

package port

import (
	"github.com/sacloud/packages-go/validate"
	"github.com/sacloud/phy-service-go/server"
)

type UpdateRequest struct {
	Id       int    `service:"-" validate:"required"`
	ServerId string `service:"-" validate:"required"`

	// ポート名称
	Nickname *string `validate:"omitempty,max=50"`
	// 有効/無効
	Enabled *bool `validate:"omitempty"`
	// ネットワーク接続設定
	Network *server.NetworkSetting `validate:"omitempty"`
}

func (req *UpdateRequest) Validate() error {
	return validate.New().Struct(req)
}