package request

import "backend-go/internal/app/model"

type GetQuestListRequest struct {
	model.Quest
	PageInfo
	Address   string
	OrderKey  string `json:"order_key" form:"order_key,default=token_id"` // 排序
	Desc      bool   `json:"desc" form:"desc,default=true"`               // 排序方式:升序false(默认)|降序true
	SearchKey string `json:"search_key" form:"search_key"`                // 搜索关键字
	Language  string `json:"-" form:"-"`                                  // 语言
}

type AddQuestRequest struct {
	Uri         string `json:"uri" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartTs     string `json:"start_ts" binding:"required"`
	EndTs       string `json:"end_ts" binding:"required"`
	Supply      string `json:"supply" binding:"required"`
}

type UpdateQuestRequest struct {
	TokenId     string `json:"token_id" binding:"required"`
	Uri         string `json:"uri" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartTs     string `json:"start_ts" binding:"required"`
	EndTs       string `json:"end_ts" binding:"required"`
	Supply      string `json:"supply" binding:"required"`
}

type UpdateRecommendRequest struct {
	TokenId   string `json:"token_id" binding:"required"`
	Recommend string `json:"recommend"` // 推荐
}

type GetCollectionQuestRequest struct {
	ID       string `json:"id" form:"id"`
	Address  string `json:"-"`
	Language string `json:"-"`
}

type GetQuestHolderRankRequest struct {
	PageInfo
}

type AddQuestV2Request struct {
	Uri         string `json:"uri" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartTs     string `json:"start_ts" binding:"required"`
	EndTs       string `json:"end_ts" binding:"required"`
	ChainID     int64  `json:"chain_id" binding:"required"`
}

type UpdateQuestV2Request struct {
	TokenId     string `json:"token_id" binding:"required"`
	Uri         string `json:"uri" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	StartTs     string `json:"start_ts" binding:"required"`
	EndTs       string `json:"end_ts" binding:"required"`
	ChainID     int64  `json:"chain_id" binding:"required"`
}
