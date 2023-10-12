package model

import "gorm.io/datatypes"

type UserChallengeLog struct {
	ID        uint           `gorm:"primarykey"`
	Address   string         `gorm:"column:address;type:varchar(44);comment:钱包地址" json:"address" form:"address"`
	TokenId   int64          `gorm:"column:token_id" json:"token_id"`
	Answer    datatypes.JSON `gorm:"column:answer" json:"answer"`
	AddTs     int64          `gorm:"column:add_ts;autoCreateTime" json:"add_ts"`
	UserScore int64          `gorm:"column:user_score" form:"user_score" json:"user_score"` // 分数
	Pass      bool           `gorm:"column:pass;default:false" json:"pass" form:"pass"`     // 状态 false 挑战未通过 true 挑战通过
	IP        string         `gorm:"column:ip" json:"ip" form:"ip"`                         // IP 地址
}
