package response

import (
	"backend-go/internal/app/model"
	"time"
)

type GetCollectionChallengeUserRes struct {
	Users []ChallengeUsers `gorm:"users" json:"users"`
	Times int64
}

type GetCollectionQuestPageResult struct {
	List       interface{} `json:"list"`
	Collection interface{} `json:"collection"`
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
}

type GetCollectionChallengeUserPageDataResult struct {
	GetCollectionChallengeUserRes
	Total    int64 `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}

type GetCollectionRes struct {
	model.Collection
	Claimed      bool   `gorm:"-" json:"claimed"`
	BadgeTokenID string `gorm:"column:badge_token_id" json:"badge_token_id"`
	BadgeChainID int64  `gorm:"column:badge_chain_id" json:"badge_chain_id"`
}

type CheckQuestInCollectionRes struct {
	IsInCollection bool `json:"is_in_collection"`
	CollectionID   int  `json:"collection_id"`
}

type GetCollectionFlashRankRes struct {
	RankList []struct {
		Rank       int64     `gorm:"rank" json:"rank"`
		Avatar     string    `gorm:"column:avatar" json:"avatar"`
		Address    string    `gorm:"address" json:"address"`
		FinishTime time.Time `gorm:"finish_time" json:"finish_time"`
	} `gorm:"-"`
	Rank       int64     `gorm:"rank" json:"rank"`
	Avatar     string    `gorm:"column:avatar" json:"avatar"`
	Address    string    `gorm:"address" json:"address"`
	FinishTime time.Time `gorm:"finish_time" json:"finish_time"`
}

type GetCollectionHighRankRes struct {
	RankList []struct {
		Rank       int64     `gorm:"rank" json:"rank"`
		Avatar     string    `gorm:"column:avatar" json:"avatar"`
		Score      int64     `gorm:"score" json:"score"`
		Address    string    `gorm:"address" json:"address"`
		FinishTime time.Time `gorm:"finish_time" json:"finish_time"`
	} `gorm:"-"`
	Rank       int64     `gorm:"rank" json:"rank"`
	Avatar     string    `gorm:"column:avatar" json:"avatar"`
	Address    string    `gorm:"address" json:"address"`
	Score      int64     `gorm:"score" json:"score"`
	FinishTime time.Time `gorm:"finish_time" json:"finish_time"`
}

type GetCollectionHolderListRes struct {
	Rank      int64     `gorm:"rank" json:"rank"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Address   string    `gorm:"address" json:"address"`
	ClaimTime time.Time `gorm:"claim_time" json:"claim_time"`
}
