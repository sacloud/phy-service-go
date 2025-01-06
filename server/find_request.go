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

package server

import (
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type FindRequest struct {
	// 電源状態
	// キャッシュされた電源状態で絞りこむ
	PowerStatus string `validate:"omitempty,power_status" meta:",options=power_status"`

	// インターネット接続状態の絞り込み
	//
	// * common - 共用グローバルネットワークを利用
	// * void - インターネット接続なし
	// * {dedicated_subnet_id} - 指定した専用グローバルネットワークを利用
	Internet string `validate:"omitempty"`

	// ローカルネットワークの接続状態の絞り込み
	// このパラメーターが複数ある場合は全てのネットワークに接続済み(AND)が対象
	//
	// * void - ローカル接続なし
	// * {private_network_id} - 指定したローカルネットワークを利用
	PrivateNetworks []string `validate:"omitempty,unique,max=5,required"`

	// タグ検索
	// このクエリーパラメーターを複数指定した場合は **すべてのタグを設定済み(AND)** のものにマッチ
	Tags []string `validate:"omitempty,unique,max=10,dive,required"`

	// フリーワード検索
	// 下記項目の **いずれか** にマッチしたものを抽出する
	//
	// * 名前(部分一致)
	// * 説明(部分一致)
	// * タグ(部分一致)
	//
	// このクエリーパラメーターを複数指定した場合は **複数の語句すべてを含む(AND)** ものにマッチ
	FreeWords []string `validate:"omitempty,unique,max=5,dive,required"`

	// 取得数
	Limit int `validate:"omitempty,min=0"`

	// 取得開始位置
	Offset int `validate:"omitempty,min=0"`

	// 並び順指定, - から始まる場合は降順指定
	//
	// * `activated` - 利用開始日順
	// * `nickname` - 名称順
	Ordering string `validate:"omitempty,ordering" meta:",options=ordering"`

	// 付加的情報の取得範囲
	IncludeFields IncludeFields `meta:"include"`
}

func (req *FindRequest) ToRequestParameter() *v1.ListServersParams {
	params := &v1.ListServersParams{}
	if req.PowerStatus != "" {
		powerStatus := v1.ListServersParamsPowerStatus(req.PowerStatus)
		params.PowerStatus = &powerStatus
	}
	if req.Internet != "" {
		params.Internet = &req.Internet
	}
	if len(req.PrivateNetworks) > 0 {
		params.PrivateNetwork = &req.PrivateNetworks
	}
	if len(req.Tags) > 0 {
		tags := req.Tags
		params.Tag = &tags
	}
	if len(req.FreeWords) > 0 {
		words := req.FreeWords
		params.FreeWord = &words
	}
	if req.Limit > 0 {
		limit := req.Limit
		params.Limit = &limit
	}
	if req.Offset > 0 {
		offset := req.Offset
		params.Offset = &offset
	}
	if req.Ordering != "" {
		order := v1.ListServersParamsOrdering(req.Ordering)
		params.Ordering = &order
	}
	return params
}
