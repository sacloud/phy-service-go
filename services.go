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

package service

import (
	"github.com/sacloud/phy-api-go"
	dedicatedsubnet "github.com/sacloud/phy-service-go/dedicated-subnet"
	privatenetwork "github.com/sacloud/phy-service-go/private-network"
	"github.com/sacloud/phy-service-go/server"
	"github.com/sacloud/phy-service-go/server/port"
	portchannel "github.com/sacloud/phy-service-go/server/port-channel"
	"github.com/sacloud/phy-service-go/service"
	"github.com/sacloud/services"
)

// Services サービス一覧を返す
func Services(client *phy.Client) []services.Service {
	return []services.Service{
		service.New(client),
		dedicatedsubnet.New(client),
		privatenetwork.New(client),
		server.New(client),
		port.New(client),
		portchannel.New(client),
	}
}
