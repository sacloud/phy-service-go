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

package portchannel

import (
	"context"

	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

// Configure ポートチャネル ボンディング設定
//
// 既存の設定が存在する場合に実行すると上書き動作(初期化)となる
func (s *Service) Configure(req *ConfigureRequest) (*v1.PortChannel, error) {
	return s.ConfigureWithContext(context.Background(), req)
}

// ConfigureWithContext ポートチャネル ボンディング設定
//
// 既存の設定が存在する場合に実行すると上書き動作(初期化)となる
func (s *Service) ConfigureWithContext(ctx context.Context, req *ConfigureRequest) (*v1.PortChannel, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	client := phy.NewServerOp(s.client)
	return client.ConfigureBonding(ctx, v1.ServerId(req.ServerId), v1.PortChannelId(req.Id), req.ToRequestParameter())
}
