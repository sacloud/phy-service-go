// Copyright 2022-2023 The sacloud/phy-service-go Authors
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
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type UpdateRequest struct {
	Id string `service:"-" validate:"required"`

	// メモ：サーバーやネットワークなどの説明
	Description *string `json:"description" validate:"omitempty"`

	// 名称：サーバーやネットワークなどの表示名
	Nickname *string `json:"nickname" validate:"omitempty"`
}

func (req *UpdateRequest) ToRequestParameter(current *v1.Service) v1.UpdateServiceParameter {
	params := v1.UpdateServiceParameter{
		Description: current.Description,
		Nickname:    current.Nickname,
	}
	if req.Description != nil {
		params.Description = req.Description
	}

	if req.Nickname != nil {
		params.Nickname = *req.Nickname
	}
	return params
}
