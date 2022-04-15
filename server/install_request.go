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
)

type InstallRequest struct {
	Id string `service:"-" validate:"required"`

	//インストールするOSイメージ
	OsImageId string `validate:"required"`

	// パスワード
	//
	// Note: 入力規則をprintasciiにしているが、本来の制限は以下の通り。
	// [0-9a-zA-Z!"#$%&'()*+,\-./:;<=>?@[\]^_`{|}~]+
	// 実装を容易にするために簡易的なバリデーションのみとしている。
	Password string `validate:"required,min=8,max=32,printascii"`

	// 手動パーティション指定
	//
	// リモートコンソールを利用し手動パーティション指定を行うか(OSが対応している場合のみ)
	ManualPartition bool
}

func (req *InstallRequest) ToRequestParameter() v1.OsInstallParameter {
	return v1.OsInstallParameter{
		ManualPartition: req.ManualPartition,
		OsImageId:       req.OsImageId,
		Password:        v1.PasswordInput(req.Password),
	}
}
