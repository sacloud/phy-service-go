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

package service

import (
	"context"

	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/services/helper"
)

func (s *Service) Update(req *UpdateRequest) (*v1.Service, error) {
	return s.UpdateWithContext(context.Background(), req)
}

func (s *Service) UpdateWithContext(ctx context.Context, req *UpdateRequest) (*v1.Service, error) {
	if err := helper.ValidateStruct(s, req); err != nil {
		return nil, err
	}
	client := phy.NewServiceOp(s.client)
	current, err := client.Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return client.Update(ctx, req.Id, req.ToRequestParameter(current))
}
