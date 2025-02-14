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

package portchannel

import (
	"github.com/sacloud/phy-api-go"
	"github.com/sacloud/phy-service-go/server"
	"github.com/sacloud/services"
	"github.com/sacloud/services/meta"
)

var _ services.Service = (*Service)(nil)

// Service provides a high-level API of for Service
type Service struct {
	client *phy.Client
}

func (s *Service) Info() *services.Info {
	return &services.Info{
		Name:           "port-channel",
		ParentServices: []string{"server"},
	}
}

func (s *Service) Operations() []services.SupportedOperation {
	return []services.SupportedOperation{
		{Name: "read", OperationType: services.OperationsRead},
		{Name: "configure", OperationType: services.OperationsUpdate},
	}
}

func (s *Service) Config() *services.Config {
	config := &services.Config{
		OptionDefs: []*meta.Option{
			{
				Key:    "bonding_type",
				Values: []string{"lacp", "static", "single"},
			},
		},
	}
	config.OptionDefs = append(config.OptionDefs, server.NetworkSettingOptions...)
	return config
}

// New returns new service instance of Service
func New(client *phy.Client) *Service {
	return &Service{client: client}
}
