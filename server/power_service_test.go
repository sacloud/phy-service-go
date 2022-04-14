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
	"testing"
	"time"

	client "github.com/sacloud/api-client-go"
	"github.com/sacloud/phy-api-go"
	"github.com/stretchr/testify/require"
)

func TestService_Power(t *testing.T) {
	fakeServer := initFakeServer()
	apiClient := &phy.Client{
		APIRootURL: fakeServer.URL,
		Options: &client.Options{
			UserAgent: "phy-service-go/test",
		},
	}
	svc := New(apiClient)

	cases := []struct {
		operation string
	}{
		{operation: "soft"},
		{operation: "on"},
		{operation: "reset"},
		{operation: "off"},
	}

	for _, tc := range cases {
		t.Run(tc.operation, func(t *testing.T) {
			require.NoError(t, svc.Power(&PowerRequest{
				Id:        serverId,
				Operation: tc.operation,

				Timeout:  time.Second,
				Interval: time.Millisecond,
			}))
		})
	}
}
