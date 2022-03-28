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

package port

import (
	"net/http/httptest"
	"testing"
	"time"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/packages-go/pointer"
	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/phy-api-go/fake"
	"github.com/sacloud/phy-api-go/fake/server"
	service "github.com/sacloud/phy-service-go"
	serverService "github.com/sacloud/phy-service-go/server"
	"github.com/stretchr/testify/require"
)

var serverId = "100000000001"
var portChannelId = 1001
var portId = 2001

func TestServerPort_CRUD_plus_L(t *testing.T) {
	fakeServer := initFakeServer()
	apiClient := &phy.Client{
		APIRootURL: fakeServer.URL,
		Options: &client.Options{
			UserAgent: service.UserAgent,
		},
	}
	svc := New(apiClient)
	var data *v1.InterfacePort

	t.Run("read without additional fields", func(t *testing.T) {
		read, err := svc.Read(&ReadRequest{
			Id:       portId,
			ServerId: serverId,
		})
		require.NoError(t, err)
		require.NotNil(t, read)

		data = read
	})

	t.Run("read return NotFoundError when server is not found", func(t *testing.T) {
		id := 9999
		read, err := svc.Read(&ReadRequest{
			Id:       id,
			ServerId: serverId,
		})
		require.Nil(t, read)
		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("list", func(t *testing.T) {
		found, err := svc.Find(&FindRequest{
			ServerId: serverId,
		})
		require.NoError(t, err)
		require.Len(t, found, 1)

		require.Equal(t, data, found[0])
	})

	t.Run("update", func(t *testing.T) {
		name := data.Nickname + "-upd"

		upd, err := svc.Update(&UpdateRequest{
			Id:       portId,
			ServerId: serverId,
			Nickname: pointer.NewString(name),
			Enabled:  pointer.NewBool(false),
			Network: &serverService.NetworkSetting{
				Mode:         "access",
				InternetType: "common_subnet",
			},
		})
		require.NoError(t, err)
		require.Equal(t, name, upd.Nickname)
		require.False(t, upd.Enabled)

		require.Equal(t, v1.InternetSubnetType("common_subnet"), upd.Internet.SubnetType)
	})

	t.Run("traffic graph", func(t *testing.T) {
		read, err := svc.TrafficGraph(&TrafficGraphRequest{
			Id:       portId,
			ServerId: serverId,
		})
		require.NoError(t, err)
		require.NotNil(t, read)
	})
}

func initFakeServer() *httptest.Server {
	raidOverallStatus := v1.RaidStatusOverallStatusOk
	fakeServer := &server.Server{
		Engine: &fake.Engine{
			Servers: []*fake.Server{
				{
					Server: &v1.Server{
						CachedPowerStatus: &v1.CachedPowerStatus{
							Status: v1.CachedPowerStatusStatusOn,
							Stored: time.Now(),
						},
						Ipv4: &v1.ServerIpv4Global{
							GatewayAddress: "192.0.2.1",
							IpAddress:      "192.0.2.11",
							NameServers:    []string{"198.51.100.1", "198.51.100.2"},
							NetworkAddress: "192.0.2.0",
							PrefixLength:   24,
							Type:           v1.ServerIpv4GlobalTypeCommonIpAddress,
						},
						LockStatus: nil,
						PortChannels: []v1.PortChannel{
							{
								BondingType:   v1.BondingTypeLacp,
								LinkSpeedType: v1.PortChannelLinkSpeedTypeN1gbe,
								Locked:        false,
								PortChannelId: portChannelId,
								Ports:         []int{portId},
							},
						},
						Ports: []v1.InterfacePort{
							{
								Enabled:             true,
								GlobalBandwidthMbps: nil,
								Internet:            nil,
								LocalBandwidthMbps:  nil,
								Mode:                nil,
								Nickname:            "server01-port01",
								PortChannelId:       portChannelId,
								PortId:              portId,
								PrivateNetworks:     nil,
							},
						},
						ServerId: serverId,
						Service: v1.ServiceQuiet{
							Activated:   time.Now(),
							Description: nil,
							Nickname:    "server01",
							ServiceId:   serverId,
							Tags:        nil,
						},
						Spec: v1.ServerSpec{
							CpuClockSpeed:         3,
							CpuCoreCount:          4,
							CpuCount:              1,
							CpuModelName:          "E3-1220 v6",
							MemorySize:            8,
							PortChannel10gbeCount: 0,
							PortChannel1gbeCount:  1,
							Storages: []v1.Storage{
								{
									BusType:     v1.StorageBusTypeSata,
									DeviceCount: 2,
									MediaType:   v1.StorageMediaTypeSsd,
									Size:        1000,
								},
							},
							TotalStorageDeviceCount: 1,
						},
						Zone: v1.Zone{
							Region: "is",
							ZoneId: 302,
						},
					},
					RaidStatus: &v1.RaidStatus{
						LogicalVolumes: []v1.RaidLogicalVolume{
							{
								PhysicalDeviceIds: []string{"0", "1"},
								RaidLevel:         "1",
								Status:            v1.RaidLogicalVolumeStatusOk,
								VolumeId:          "0",
							},
						},
						Monitored:     time.Now(),
						OverallStatus: &raidOverallStatus,
						PhysicalDevices: []v1.RaidPhysicalDevice{
							{
								DeviceId: "0",
								Slot:     0,
								Status:   v1.RaidPhysicalDeviceStatusOk,
							},
							{
								DeviceId: "1",
								Slot:     1,
								Status:   v1.RaidPhysicalDeviceStatusOk,
							},
						},
					},
					OSImages: []*v1.OsImage{
						{
							ManualPartition: true,
							Name:            "Usacloud Linux",
							OsImageId:       "usacloud",
							RequirePassword: true,
							SuperuserName:   "root",
						},
					},
					PowerStatus: &v1.ServerPowerStatus{
						Status: v1.ServerPowerStatusStatusOn,
					},
					TrafficGraph: &v1.TrafficGraph{
						Receive: []v1.TrafficGraphData{
							{
								Timestamp: time.Now(),
								Value:     1,
							},
						},
						Transmit: []v1.TrafficGraphData{
							{
								Timestamp: time.Now(),
								Value:     1,
							},
						},
					},
				},
			},
		},
	}
	return httptest.NewServer(fakeServer.Handler())
}
