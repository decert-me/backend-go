package model

import "gorm.io/gorm"

type UserChallengeClaim struct {
	gorm.Model
	Address string `gorm:"column:address;type:varchar(44);comment:钱包地址" json:"address" form:"address"`
	TokenId string `gorm:"column:token_id;type:varchar(100)" json:"token_id"`
	Status  uint8  `gorm:"column:status;default:1" json:"status"` // 状态 1: 待空投 2: 成功空投
	ChainID int64  `gorm:"column:chain_id;comment:链ID;default:137" json:"chain_id"`
}
