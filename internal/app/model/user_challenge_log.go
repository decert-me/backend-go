package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserChallengeLog struct {
	gorm.Model
	Address     string         `gorm:"column:address;type:varchar(44);comment:钱包地址;index:address_tokenId" json:"address" form:"address"`
	TokenId     string         `gorm:"column:token_id;index:address_tokenId;type:varchar(100)" json:"token_id"`
	Answer      datatypes.JSON `gorm:"column:answer" json:"answer"`
	AddTs       int64          `gorm:"column:add_ts;autoCreateTime" json:"add_ts"`
	UserScore   int64          `gorm:"column:user_score" form:"user_score" json:"user_score"`                                      // 分数
	Pass        bool           `gorm:"column:pass;default:false" json:"pass" form:"pass"`                                          // 状态 false 挑战未通过 true 挑战通过
	IP          string         `gorm:"column:ip" json:"ip" form:"ip"`                                                              // IP 地址
	IsOpenQuest bool           `gorm:"column:is_open_quest;default:false;comment:是否开放题" json:"is_open_quest" form:"is_open_quest"` // 是否开放题
}
