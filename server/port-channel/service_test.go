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
	"net/http/httptest"
	"testing"
	"time"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/phy-api-go/fake"
	"github.com/sacloud/phy-api-go/fake/server"
	serverService "github.com/sacloud/phy-service-go/server"
	"github.com/stretchr/testify/require"
)

var serverId = "100000000001"
var portChannelId = 1001

func TestServerPortChannel_CRUD_plus_L(t *testing.T) {
	fakeServer := initFakeServer()
	apiClient := &phy.Client{
		APIRootURL: fakeServer.URL,
		Options: &client.Options{
			UserAgent: "phy-service-go/test",
		},
	}
	svc := New(apiClient)
	var portChannel *v1.PortChannel

	t.Run("read", func(t *testing.T) {
		read, err := svc.Read(&ReadRequest{
			Id:       portChannelId,
			ServerId: serverId,
		})
		require.NoError(t, err)
		require.NotNil(t, read)

		portChannel = read
	})

	t.Run("read return NotFoundError when port-channel is not found", func(t *testing.T) {
		id := 999999
		read, err := svc.Read(&ReadRequest{
			Id:       id,
			ServerId: serverId,
		})
		require.Nil(t, read)
		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("configure", func(t *testing.T) {
		updated, err := svc.Configure(&ConfigureRequest{
			Id:          portChannelId,
			ServerId:    serverId,
			BondingType: "lacp",
			PortSettings: []*PortSetting{
				{
					Nickname: "bond01",
					Network: serverService.NetworkSetting{
						Mode:         "access",
						InternetType: "common_subnet",
					},
				},
			},
		})

		require.NoError(t, err)
		require.NotEqualValues(t, portChannel.Ports, updated.Ports)
		require.Equal(t, v1.BondingType("lacp"), updated.BondingType)
		require.NotEmpty(t, updated.ConfiguredPorts)
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
								Ports:         []int{2001},
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
								PortChannelId:       1001,
								PortId:              2001,
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
