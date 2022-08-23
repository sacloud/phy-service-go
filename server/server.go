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

package server

import (
	"context"

	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type Server struct {
	v1.Server
	*v1.RaidStatus
	*v1.ServerPowerStatus
	OsImages []*v1.OsImage
}

type IncludeFields struct {
	// RAID状態(キャッシュ)を含むか
	//
	// RefreshedRaidStatusがtrueの場合はそちらが優先される
	CachedRaidStatus bool
	// RAID状態(最新)を含むか
	RefreshedRaidStatus bool

	// 電源状態を含むか
	PowerStatus bool
	// インストール可能OS一覧情報を含むか
	OSImages bool
}

func (s *Service) fetchAdditionalInfo(ctx context.Context, server *Server, fields IncludeFields) error {
	client := phy.NewServerOp(s.client)
	if fields.CachedRaidStatus || fields.RefreshedRaidStatus {
		status, err := client.ReadRAIDStatus(ctx, server.ServerId, fields.RefreshedRaidStatus)
		if err != nil {
			return err
		}
		server.RaidStatus = status
	}
	if fields.PowerStatus {
		status, err := client.ReadPowerStatus(ctx, server.ServerId)

		if err != nil {
			return err
		}
		server.ServerPowerStatus = status
	}
	if fields.OSImages {
		images, err := client.ListOSImages(ctx, server.ServerId)
		if err != nil {
			return err
		}
		server.OsImages = images
	}
	return nil
}
