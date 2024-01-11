package model

import (
	"gorm.io/datatypes"
)

type UserChallenges struct {
	ID         uint           `gorm:"primarykey"`
	Address    string         `gorm:"column:address;type:varchar(44);index:,unique,composite:challenges_address_tokenId;comment:钱包地址" json:"address" form:"address"`
	TokenId    string         `gorm:"column:token_id;index:,unique,composite:challenges_address_tokenId;type:varchar(100)" json:"tokenId"`
	Status     uint8          `gorm:"column:status;default:0;size:30;comment:0 进行中 1 等待验证 2 成功;" json:"status" form:"status"` // 0:进行中;1:等待验证;2:成功;
	Content    datatypes.JSON `gorm:"column:content" json:"content"`
	AddTs      int64          `gorm:"column:add_ts;autoCreateTime" json:"addTs"`
	Claimed    bool           `gorm:"column:claimed;default:false" json:"claimed"`
	UpdateTs   int64          `gorm:"column:update_ts;autoUpdateTime" json:"updateTs"`
	UserScore  int64          `gorm:"column:user_score" form:"user_score" json:"user_score"` // 分数
	ClaimTs    int64          `gorm:"column:claim_ts;" json:"claimTs"`
	NFTAddress string         `gorm:"column:nft_address" json:"nft_address" form:"nft_address"`
}
