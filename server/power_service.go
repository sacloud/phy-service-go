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
	"fmt"

	"github.com/sacloud/packages-go/wait"
	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/services/helper"
)

func (s *Service) Power(req *PowerRequest) error {
	return s.PowerWithContext(context.Background(), req)
}

func (s *Service) PowerWithContext(ctx context.Context, req *PowerRequest) error {
	if err := helper.ValidateStruct(s, req); err != nil {
		return err
	}
	client := phy.NewServerOp(s.client)
	if err := client.PowerControl(ctx, v1.ServerId(req.Id), v1.ServerPowerOperations(req.Operation)); err != nil {
		return err
	}

	if req.NoWait {
		return nil
	}
	return s.waitAfterPowerControl(ctx, req, client)
}

func (s *Service) waitAfterPowerControl(ctx context.Context, req *PowerRequest, client phy.ServerAPI) error {
	desiredPowerState := v1.ServerPowerStatusStatusOn
	switch req.Operation {
	case "on", "reset":
		desiredPowerState = v1.ServerPowerStatusStatusOn
	case "soft", "off":
		desiredPowerState = v1.ServerPowerStatusStatusOff
	default:
		panic(fmt.Sprintf("unexpected power operation: %s", req.Operation))
	}

	waiter := wait.PollingWaiter{
		ReadFunc: func() (interface{}, error) {
			return client.ReadPowerStatus(ctx, v1.ServerId(req.Id))
		},
		StateCheckFunc: func(target interface{}) (bool, error) {
			if state, ok := target.(*v1.ServerPowerStatus); ok {
				return state.Status == desiredPowerState, nil
			}
			return false, nil
		},
		Timeout:  req.Timeout,
		Interval: req.Interval,
	}

	_, err := waiter.WaitForState(ctx)
	return err
}
