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

package service

import (
	"fmt"
	"runtime"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/phy-api-go"
)

// UserAgent APIリクエスト時のユーザーエージェント
var UserAgent = fmt.Sprintf(
	"phy-service-go/v%s (%s/%s; +https://github.com/sacloud/phy-service-go) %s",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
	phy.UserAgent,
)

func NewClient() *phy.Client {
	return &phy.Client{
		Options: &client.Options{
			UserAgent: UserAgent,
		},
	}
}
