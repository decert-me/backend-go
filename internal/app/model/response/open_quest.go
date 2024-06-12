package response

import (
	"time"

	"gorm.io/datatypes"
)

type UserOpenQuestJsonElements struct {
	ID                    uint           `gorm:"primarykey"`
	UUID                  string         `gorm:"column:uuid" json:"uuid"`
	Address               string         `gorm:"column:address;type:varchar(44);comment:钱包地址;index:address_tokenId" json:"address" form:"address"`
	TokenId               string         `gorm:"column:token_id;index:address_tokenId" json:"token_id"`
	OpenQuestReviewStatus uint8          `gorm:"column:open_quest_review_status;default:0;comment:评阅开放题状态 1 未审核 2 已审核" json:"open_quest_review_status" form:"open_quest_review_status"` // // 评阅开放题状态 1 未审核 2 已审核
	OpenQuestReviewTime   string         `gorm:"column:open_quest_review_time;comment:评阅开放题时间" json:"open_quest_review_time" form:"open_quest_review_time"`
	UpdatedAt             time.Time      `gorm:"column:updated_at" json:"updated_at"`
	Index                 int            `gorm:"column:index" json:"index"`
	Title                 string         `gorm:"column:title" json:"title"`
	Answer                datatypes.JSON `gorm:"column:answer" json:"answer"`
	ChallengeTitle        string         `gorm:"column:challenge_title" json:"challenge_title"`
	Score                 int64          `gorm:"column:score" json:"score"`
	Correct               bool           `gorm:"column:correct" json:"correct"`
	UserAnswer            datatypes.JSON `gorm:"column:user_answer" json:"-"`
	MetaData              datatypes.JSON `gorm:"column:meta_data" json:"-"`                                // 元数据
	QuestData             datatypes.JSON `gorm:"column:quest_data" json:"-"`                               // 元数据
	PassScore             int64          `gorm:"column:pass_score" form:"pass_score" json:"pass_score"`    // 通过分数
	TotalScore            int64          `gorm:"column:total_score" form:"total_score" json:"total_score"` // 总分
	UserScore             int64          `gorm:"column:user_score" form:"user_score" json:"user_score"`    // 用户分数
	NickName              *string        `gorm:"column:nickname;type:varchar(200);default:''" json:"nickname" form:"nickname"`
}

type UserOpenQuestJsonElementsV2 struct {
	UUID           string    `gorm:"column:uuid" json:"uuid"`
	TokenId        string    `gorm:"column:token_id;index:address_tokenId" json:"token_id"`
	Index          int       `gorm:"column:index" json:"index"`
	Title          string    `gorm:"column:title" json:"title"`
	ChallengeTitle string    `gorm:"column:challenge_title" json:"challenge_title"`
	ToReviewCount  int64     `gorm:"-" json:"to_review_count"`  // 待评分数量
	ReviewedCount  int64     `gorm:"-" json:"reviewed_count"`   // 已评分数量
	LastSummitTime time.Time `gorm:"-" json:"last_sumbit_time"` // 最新提交时间
	LastReviewTime time.Time `gorm:"-" json:"last_review_time"` // 上次评分时间
	Addts          int64     `gorm:"column:add_ts" json:"add_ts"`
}
