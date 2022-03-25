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

package service

import (
	"github.com/sacloud/packages-go/validate"
	v1 "github.com/sacloud/phy-api-go/apis/v1"
)

type FindRequest struct {
	// サービス種類検索
	//
	// * server
	// * dedicated_subnet
	// * private_network
	// * firewall
	// * load_balancer
	ProductCategory string `validate:"omitempty,oneof=server dedicated_subnet private_network firewall load_balancer"`

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
	Ordering string `validate:"omitempty,oneof=activated -activated nickname -nickname"`
}

func (req *FindRequest) Validate() error {
	return validate.New().Struct(req)
}

func (req *FindRequest) ToRequestParameter() *v1.ListServicesParams {
	params := &v1.ListServicesParams{}
	if req.ProductCategory != "" {
		category := v1.ListServicesParamsProductCategory(req.ProductCategory)
		params.ProductCategory = &category
	}
	if len(req.Tags) > 0 {
		tags := v1.TagFilter(req.Tags)
		params.Tag = &tags
	}
	if len(req.FreeWords) > 0 {
		words := v1.FreeWordFilter(req.FreeWords)
		params.FreeWord = &words
	}
	if req.Limit > 0 {
		limit := v1.Limit(req.Limit)
		params.Limit = &limit
	}
	if req.Offset > 0 {
		offset := v1.Offset(req.Offset)
		params.Offset = &offset
	}
	if req.Ordering != "" {
		order := v1.ListServicesParamsOrdering(req.Ordering)
		params.Ordering = &order
	}
	return params
}
