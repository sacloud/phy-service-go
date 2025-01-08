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
	"github.com/sacloud/services"
	"github.com/sacloud/services/meta"
)

// Service provides a high-level API of for Service
type Service struct {
	client *phy.Client
}

var _ services.Service = (*Service)(nil)

// New returns new service instance of Service
func New(client *phy.Client) *Service {
	return &Service{client: client}
}

func (s *Service) Info() *services.Info {
	return &services.Info{
		Name: "service",
	}
}

func (s *Service) Operations() []services.SupportedOperation {
	return []services.SupportedOperation{
		{Name: "find", OperationType: services.OperationsList},
		{Name: "read", OperationType: services.OperationsRead},
		{Name: "update", OperationType: services.OperationsUpdate},
	}
}

func (s *Service) Config() *services.Config {
	return &services.Config{
		OptionDefs: []*meta.Option{
			{
				Key:    "product_category",
				Values: []string{"server", "dedicated_subnet", "private_network", "firewall", "load_balancer"},
			},
			{
				Key:    "ordering",
				Values: []string{"activated", "-activated", "nickname", "-nickname"},
			},
		},
	}
}
