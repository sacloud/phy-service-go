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

package server

import (
	"context"

	"github.com/sacloud/phy-api-go"
	"github.com/sacloud/services/helper"
)

func (s *Service) Install(req *InstallRequest) error {
	return s.InstallWithContext(context.Background(), req)
}

func (s *Service) InstallWithContext(ctx context.Context, req *InstallRequest) error {
	if err := helper.ValidateStruct(s, req); err != nil {
		return err
	}
	client := phy.NewServerOp(s.client)
	return client.OSInstall(ctx, req.Id, req.ToRequestParameter())
}
