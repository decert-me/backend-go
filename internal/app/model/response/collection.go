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
	Claimed bool `gorm:"-" json:"claimed"`
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
