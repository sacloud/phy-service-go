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
	"context"

	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/services/helper"
)

func (s *Service) Update(req *UpdateRequest) (*v1.InterfacePort, error) {
	return s.UpdateWithContext(context.Background(), req)
}

func (s *Service) UpdateWithContext(ctx context.Context, req *UpdateRequest) (*v1.InterfacePort, error) {
	if err := helper.ValidateStruct(s, req); err != nil {
		return nil, err
	}
	client := phy.NewServerOp(s.client)

	var updated *v1.InterfacePort

	if req.Nickname != nil {
		upd, err := client.UpdatePort(ctx,
			req.ServerId, req.Id,
			v1.UpdateServerPortParameter{Nickname: *req.Nickname})
		if err != nil {
			return nil, err
		}
		updated = upd
	}

	if req.Enabled != nil {
		upd, err := client.EnablePort(ctx,
			req.ServerId, req.Id,
			*req.Enabled)
		if err != nil {
			return nil, err
		}
		updated = upd
	}

	if req.Network != nil {
		upd, err := client.AssignNetwork(ctx,
			req.ServerId, req.Id,
			req.Network.ToRequestParameter())
		if err != nil {
			return nil, err
		}
		updated = upd
	}

	return updated, nil
}
