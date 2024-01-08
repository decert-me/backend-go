package response

import (
	"backend-go/internal/app/model"
	"gorm.io/datatypes"
	"time"
)

type GetQuestListRes struct {
	model.Quest
	Claimed               bool        `gorm:"claimed" json:"claimed"`
	CollectionCount       int64       `gorm:"-" json:"collection_count"`
	EstimateTime          interface{} `gorm:"-" json:"estimate_time"` // 预估时间/min
	AuthorInfo            model.Users `gorm:"-" json:"author_info"`
	OpenQuestReviewStatus uint8       `gorm:"column:open_quest_review_status" json:"open_quest_review_status"` // 评阅开放题状态 1 未审核 2 已审核
	Claimable             bool        `gorm:"claimable" json:"claimable"`                                      // 是否可领取
}

type ChallengeUsers struct {
	ID          uint    `gorm:"primarykey" json:"-"`
	Address     string  `gorm:"column:address;type:varchar(44);UNIQUE;comment:钱包地址" json:"address" form:"address"`
	NickName    *string `gorm:"column:nickname;type:varchar(200);default:''" json:"nickname" form:"nickname"`
	Avatar      *string `gorm:"column:avatar;type:varchar(200);comment:用户头像;default:''" json:"avatar" form:"avatar"`
	Description *string `gorm:"column:description;type:varchar(100);comment:自我介绍;default:''" json:"description" form:"description"`
}

type GetQuestChallengeUserRes struct {
	Users []ChallengeUsers `gorm:"users" json:"users"`
	Times int64
}

type GetQuestRes struct {
	model.Quest
	Claimed               bool           `gorm:"claimed" json:"claimed"`
	UserScore             int64          `gorm:"user_score" json:"user_score"`
	NFTAddress            string         `gorm:"column:nft_address" json:"nft_address"`
	Answer                datatypes.JSON `gorm:"column:answer" json:"answer"`
	OpenQuestReviewStatus uint8          `gorm:"column:open_quest_review_status" json:"open_quest_review_status"` // 评阅开放题状态 1 未审核 2 已审核
	Answers               []string       `gorm:"-" json:"answers"`
}

type GetQuestFlashListRes struct {
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
	//Total      int64     `gorm:"total" json:"total"`
}

type GetQuestHighScoreListRes struct {
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

type GetQuestHolderListRes struct {
	Rank      int64     `gorm:"rank" json:"rank"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Address   string    `gorm:"address" json:"address"`
	ClaimTime time.Time `gorm:"claim_time" json:"claim_time"`
}

type GetAddressHighScore struct {
	Address string         `gorm:"address"`
	TokenId int64          `gorm:"token_id"`
	Score   int64          `gorm:"score"`
	Answer  datatypes.JSON `gorm:"answer"`
}
