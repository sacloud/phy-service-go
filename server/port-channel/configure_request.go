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
	"github.com/sacloud/packages-go/validate"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type ConfigureRequest struct {
	Id       int    `service:"-" validate:"required"`
	ServerId string `service:"-" validate:"required"`

	// ボンディング方式指定
	//
	// * lacp - LACP
	// * static - static link aggregation
	// * single - ボンディングなし(単体構成)
	BondingType string `validate:"required,oneof=lacp static single"`

	// 作成するポート名称の指定
	//
	// * nil: 自動設定
	// * 1要素: ボンディング構成
	// * 2要素: ボンディングなし
	PortNicknames *[]string `validate:"omitempty,min=1,max=2,dive,required"`
}

func (req *ConfigureRequest) Validate() error {
	return validate.New().Struct(req)
}

func (req *ConfigureRequest) ToRequestParameter() v1.ConfigureBondingParameter {
	return v1.ConfigureBondingParameter{
		BondingType:   v1.BondingType(req.BondingType),
		PortNicknames: req.PortNicknames,
	}
}
