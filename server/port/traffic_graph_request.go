// Copyright 2022-2025 The sacloud/phy-service-go Authors
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
	"time"

	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type TrafficGraphRequest struct {
	Id       int    `service:"-" validate:"required"`
	ServerId string `service:"-" validate:"required"`

	// 取得範囲始点(過去31日前まで,未指定時は7日前)
	Since time.Time `validate:"omitempty"`
	// 取得範囲終点(未指定時は現在時刻)
	Until time.Time `validate:"omitempty"`
	// データポイント間隔(秒)
	Step int `validate:"omitempty,step" meta:",options=step"`
}

func (req *TrafficGraphRequest) ToRequestParameter() v1.ReadServerTrafficByPortParams {
	params := v1.ReadServerTrafficByPortParams{}
	if !req.Since.IsZero() {
		params.Since = &req.Since
	}
	if !req.Until.IsZero() {
		params.Until = &req.Until
	}
	if req.Step > 0 {
		v := v1.ReadServerTrafficByPortParamsStep(req.Step)
		params.Step = &v
	}
	return params
}
