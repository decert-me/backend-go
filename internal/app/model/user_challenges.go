package model

import (
	"gorm.io/datatypes"
)

type UserChallenges struct {
	ID       uint           `gorm:"primarykey"`
	Address  string         `gorm:"column:address;type:char(42);index:challenges_address_tokenId,UNIQUE;comment:钱包地址" json:"address" form:"address"`
	TokenId  int64          `gorm:"column:token_id;index:challenges_address_tokenId,UNIQUE;" json:"tokenId"`
	Status   uint8          `gorm:"column:status;default:0;size:30;comment:0 进行中 1 等待验证 2 成功;" json:"status" form:"status"` // 0:进行中;1:等待验证;2:成功;
	Content  datatypes.JSON `gorm:"column:content" json:"content"`
	AddTs    int64          `gorm:"column:add_ts;autoCreateTime" json:"addTs"`
	Claimed  bool           `gorm:"column:claimed;default:false" json:"claimed"`
	UpdateTs int64          `gorm:"column:update_ts;autoUpdateTime" json:"updateTs"`
	ClaimTs  int64          `gorm:"column:claim_ts;" json:"claimTs"`
}
