package model

import "gorm.io/gorm"

type UserChallengeClaim struct {
	gorm.Model
	Address string `gorm:"column:address;type:char(44);comment:钱包地址" json:"address" form:"address"`
	TokenId int64  `gorm:"column:token_id" json:"token_id"`
	Status  uint8  `gorm:"column:status;default:1" json:"status"` // 状态 1: 待空投 2: 成功空投
}
