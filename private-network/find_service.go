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

package privatenetwork

import (
	"context"

	"github.com/sacloud/phy-api-go"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

func (s *Service) Find(req *FindRequest) ([]*v1.PrivateNetwork, error) {
	return s.FindWithContext(context.Background(), req)
}

func (s *Service) FindWithContext(ctx context.Context, req *FindRequest) ([]*v1.PrivateNetwork, error) {
	if req == nil {
		req = &FindRequest{}
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}

	client := phy.NewPrivateNetworkOp(s.client)
	found, err := client.List(ctx, req.ToRequestParameter())
	if err != nil {
		return nil, err
	}

	var results []*v1.PrivateNetwork
	for i := range found.PrivateNetworks {
		results = append(results, &found.PrivateNetworks[i])
	}
	return results, nil
}
