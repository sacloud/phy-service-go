// Copyright 2022-2023 The sacloud/phy-service-go Authors
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

package server

import (
	"time"
)

type PowerRequest struct {
	Id string `service:"-" validate:"required"`

	// 電源操作内容
	Operation string `validate:"required,power_operation" meta:",options=power_operation"`

	// 電源操作後に希望の状態になるまで待つか
	NoWait bool

	// 待ち処理のタイムアウト
	Timeout time.Duration `validate:"omitempty,min=0"`

	// 待ち処理の処理間隔
	Interval time.Duration `validate:"omitempty,min=0"`
}
