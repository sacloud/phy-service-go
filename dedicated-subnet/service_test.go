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

package dedicatedsubnet

import (
	"net/http/httptest"
	"testing"
	"time"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/phy-api-go/fake"
	"github.com/sacloud/phy-api-go/fake/server"
	"github.com/stretchr/testify/require"
)

var dedicatedSubnetId = "100000000001"

func TestAccount_CRUD_plus_L(t *testing.T) {
	fakeServer := initFakeServer()
	apiClient := &phy.Client{
		APIRootURL: fakeServer.URL,
		Options: &client.Options{
			UserAgent: "phy-service-go/test",
		},
	}
	svc := New(apiClient)
	var data *v1.DedicatedSubnet

	t.Run("read", func(t *testing.T) {
		read, err := svc.Read(&ReadRequest{
			Id: dedicatedSubnetId,
		})
		require.NoError(t, err)
		require.NotNil(t, read)

		data = read
	})

	t.Run("read return NotFoundError when account is not found", func(t *testing.T) {
		id := "not-exists-account-id"
		read, err := svc.Read(&ReadRequest{
			Id: id,
		})
		require.Nil(t, read)
		require.Error(t, err)
		require.True(t, v1.IsError404(err))
	})

	t.Run("list", func(t *testing.T) {
		found, err := svc.Find(&FindRequest{})
		require.NoError(t, err)
		require.Len(t, found, 1)

		require.Equal(t, data, found[0])
	})
}

func initFakeServer() *httptest.Server {
	fakeServer := &server.Server{
		Engine: &fake.Engine{
			DedicatedSubnets: []*v1.DedicatedSubnet{
				{
					ConfigStatus:      v1.DedicatedSubnetConfigStatusOperational,
					DedicatedSubnetId: dedicatedSubnetId,
					Firewall:          nil,
					Ipv4: v1.Ipv4{
						BroadcastAddress:    "192.0.2.239",
						GatewayAddress:      "192.0.2.225",
						NetworkAddress:      "192.0.2.224",
						PrefixLength:        28,
						SpecialUseAddresses: nil,
					},
					LoadBalancer: nil,
					ServerCount:  1,
					Service: v1.ServiceQuiet{
						Activated: time.Now(),
						Nickname:  "global-network01",
						ServiceId: "100000000001",
						Tags:      nil,
					},
					Zone: v1.Zone{
						Region: "is",
						ZoneId: 302,
					},
				},
			},
		},
	}
	return httptest.NewServer(fakeServer.Handler())
}
