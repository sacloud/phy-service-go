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
	v1 "github.com/sacloud/phy-api-go/apis/v1"
	"github.com/sacloud/services/meta"
)

var NetworkSettingOptions = []*meta.Option{
	{
		Key:    "network_mode",
		Values: []string{"access", "trunk"},
	},
	{
		Key:    "internet_type",
		Values: []string{"", "common_subnet", "dedicated_subnet"},
	},
}

type NetworkSetting struct {
	// ポートモード
	//
	// * access - アクセスポート
	// * trunk - トランクポート
	Mode string `validate:"required,network_mode" meta:",options=network_mode"`

	// インターネット接続タイプ
	// * 空 - インターネット接続なし
	// * common_subnet - 共用グローバルネットワーク利用
	// * dedicated_subnet - 専用グローバルネットワーク利用
	InternetType string `validate:"omitempty,internet_type" meta:",options=internet_type"`

	// 専用グローバルネットワークのサービスコード
	//
	//InternetTypeが`dedicated_subnet`の場合に必須
	DedicatedSubnetId string `validate:"required_if=InternetType dedicated_subnet"`

	// 	接続先ローカルネットワークの配列
	PrivateNetworkIds []string `validate:"omitempty,dive,required"`
}

func (n *NetworkSetting) ToRequestParameter() v1.AssignNetworkParameter {
	params := v1.AssignNetworkParameter{
		Mode: v1.AssignNetworkParameterMode(n.Mode),
	}

	if n.InternetType != "" {
		v := v1.AssignNetworkParameterInternetType(n.InternetType)
		params.InternetType = &v
	}

	if n.DedicatedSubnetId != "" {
		params.DedicatedSubnetId = &n.DedicatedSubnetId
	}

	if len(n.PrivateNetworkIds) > 0 {
		params.PrivateNetworkIds = &n.PrivateNetworkIds
	}

	return params
}
