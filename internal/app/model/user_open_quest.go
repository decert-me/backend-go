package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type UserOpenQuest struct {
	gorm.Model
	Address               string         `gorm:"column:address;type:char(42);comment:钱包地址" json:"address" form:"address"`
	TokenId               int64          `gorm:"column:token_id" json:"token_id"`
	Answer                datatypes.JSON `gorm:"column:answer" json:"answer"`
	OpenQuestScore        int64          `gorm:"column:open_quest_score;default:0;comment:开放题分数" json:"open_quest_score" form:"open_quest_score"`                                       // 开放题分数
	OpenQuestReviewStatus uint8          `gorm:"column:open_quest_review_status;default:0;comment:评阅开放题状态 1 未审核 2 已审核" json:"open_quest_review_status" form:"open_quest_review_status"` // // 评阅开放题状态 1 未审核 2 已审核
	OpenQuestReviewTime   time.Time      `gorm:"column:open_quest_review_time;comment:评阅开放题时间" json:"open_quest_review_time" form:"open_quest_review_time"`
}
