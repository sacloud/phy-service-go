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
	"fmt"

	"github.com/sacloud/phy-api-go"
	"github.com/sacloud/services/helper"
)

// Configure ポートチャネル ボンディング設定
//
// 既存の設定が存在する場合に実行すると上書き動作(初期化)となる
func (s *Service) Configure(req *ConfigureRequest) (*ConfiguredPortChannel, error) {
	return s.ConfigureWithContext(context.Background(), req)
}

// ConfigureWithContext ポートチャネル ボンディング設定
//
// 既存の設定が存在する場合に実行すると上書き動作(初期化)となる
func (s *Service) ConfigureWithContext(ctx context.Context, req *ConfigureRequest) (*ConfiguredPortChannel, error) {
	if err := helper.ValidateStruct(s, req); err != nil {
		return nil, err
	}
	client := phy.NewServerOp(s.client)
	portChannel, err := client.ConfigureBonding(ctx,
		req.ServerId, req.Id,
		req.ConfigureBondingParameter(),
	)
	if err != nil {
		return nil, err
	}

	configured := &ConfiguredPortChannel{
		PortChannel: *portChannel,
	}

	if len(portChannel.Ports) != len(req.PortSettings) {
		return configured, fmt.Errorf("invalid port settings: %#+v", req.PortSettings)
	}

	for i := range portChannel.Ports {
		portId := configured.Ports[i]
		setting := req.PortSettings[i]

		assigned, err := client.AssignNetwork(ctx,
			req.ServerId, portId,
			setting.Network.ToRequestParameter(),
		)
		if err != nil {
			return configured, err
		}
		configured.ConfiguredPorts = append(configured.ConfiguredPorts, assigned)
	}
	return configured, nil
}
