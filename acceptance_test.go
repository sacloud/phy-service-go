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

//go:build acctest
// +build acctest

package service

import (
	"testing"
	"time"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	dedicatedsubnet "github.com/sacloud/phy-service-go/dedicated-subnet"
	privatenetwork "github.com/sacloud/phy-service-go/private-network"
	"github.com/sacloud/phy-service-go/server"
	"github.com/sacloud/phy-service-go/server/port"
	portchannel "github.com/sacloud/phy-service-go/server/port-channel"
	"github.com/sacloud/phy-service-go/service"
	"github.com/stretchr/testify/require"
)

var serverPowerControlWaitDuration = 1 * time.Minute

func TestAccServer(t *testing.T) {
	client := testClient()

	serviceService := service.New(client)
	serverService := server.New(client)
	portService := port.New(client)
	portChannelService := portchannel.New(client)

	var serverId string
	var readServer *server.Server

	t.Run("get server id from service list", func(t *testing.T) {
		// サービスリストからサーバ一覧に絞って1件だけ取得、取得したリソースを以降のテスト対象とする
		services, err := serviceService.Find(&service.FindRequest{
			ProductCategory: "server",
			Limit:           1,
		})

		require.NoError(t, err)

		if len(services) > 0 {
			serverId = services[0].ServiceId
		}
	})

	if serverId == "" {
		t.Skip("target servers not found")
	}

	t.Run("read server", func(t *testing.T) {
		read, err := serverService.Read(&server.ReadRequest{
			Id: serverId,
			IncludeFields: server.IncludeFields{
				CachedRaidStatus:    false,
				RefreshedRaidStatus: false,
				PowerStatus:         true,
				OSImages:            true,
			},
		})
		require.NoError(t, err)

		require.NotEmpty(t, read)
		require.NotEmpty(t, read.ServerPowerStatus)
		require.NotEmpty(t, read.OsImages)

		readServer = read
	})

	t.Run("read server RAID status", func(t *testing.T) {
		if readServer.ServerPowerStatus.Status != v1.ServerPowerStatusStatusOn {
			t.Skip("server power state is not 'on'")
		}

		read, err := serverService.Read(&server.ReadRequest{
			Id: serverId,
			IncludeFields: server.IncludeFields{
				CachedRaidStatus: true,
				//RefreshedRaidStatus: true,
			},
		})
		require.NoError(t, err)

		require.NotEmpty(t, read)
		require.NotEmpty(t, read.RaidStatus)

		readServer = read
	})

	t.Run("list ports", func(t *testing.T) {
		ports, err := portService.Find(&port.FindRequest{
			ServerId: serverId,
		})
		require.NoError(t, err)
		require.NotEmpty(t, ports)
	})

	t.Run("list port channels", func(t *testing.T) {
		pc, err := portChannelService.Read(&portchannel.ReadRequest{
			ServerId: serverId,
			Id:       readServer.PortChannels[0].PortChannelId,
		})
		require.NoError(t, err)
		require.NotNil(t, pc)
	})
}

func TestAccDedicatedSubnet(t *testing.T) {
	client := testClient()

	serviceService := service.New(client)
	dedicatedSubnetService := dedicatedsubnet.New(client)

	var subnetId string

	t.Run("get dedicated-subnet id from service list", func(t *testing.T) {
		// サービスリストから1件だけ取得、取得したリソースを以降のテスト対象とする
		services, err := serviceService.Find(&service.FindRequest{
			ProductCategory: "dedicated_subnet",
			Limit:           1,
		})

		require.NoError(t, err)

		if len(services) > 0 {
			subnetId = services[0].ServiceId
		}
	})

	if subnetId == "" {
		t.Skip("target dedicated-subnet not found")
	}

	t.Run("list", func(t *testing.T) {
		subnets, err := dedicatedSubnetService.Find(&dedicatedsubnet.FindRequest{})
		require.NoError(t, err)
		require.NotEmpty(t, subnets)
	})

	t.Run("read", func(t *testing.T) {
		subnet, err := dedicatedSubnetService.Read(&dedicatedsubnet.ReadRequest{
			Id:      subnetId,
			Refresh: pointer.NewBool(true),
		})
		require.NoError(t, err)
		require.NotEmpty(t, subnet)
	})
}

func TestAccPrivateNetwork(t *testing.T) {
	client := testClient()

	serviceService := service.New(client)
	privateNetworkService := privatenetwork.New(client)

	var networkId string

	t.Run("get private-network id from service list", func(t *testing.T) {
		// サービスリストから1件だけ取得、取得したリソースを以降のテスト対象とする
		services, err := serviceService.Find(&service.FindRequest{
			ProductCategory: "private_network",
			Limit:           1,
		})

		require.NoError(t, err)

		if len(services) > 0 {
			networkId = services[0].ServiceId
		}
	})

	if networkId == "" {
		t.Skip("target private-network not found")
	}

	t.Run("list", func(t *testing.T) {
		subnets, err := privateNetworkService.Find(&privatenetwork.FindRequest{})
		require.NoError(t, err)
		require.NotEmpty(t, subnets)
	})

	t.Run("read", func(t *testing.T) {
		subnet, err := privateNetworkService.Read(&privatenetwork.ReadRequest{
			Id: networkId,
		})
		require.NoError(t, err)
		require.NotEmpty(t, subnet)
	})
}

func testClient() *phy.Client {
	return &phy.Client{
		Options: &client.Options{
			UserAgent: "phy-service-go/test",
		},
	}
}
