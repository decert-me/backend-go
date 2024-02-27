package response

import (
	"gorm.io/datatypes"
	"time"
)

type UserOpenQuestJsonElements struct {
	ID                    uint           `gorm:"primarykey"`
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
}
