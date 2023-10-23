package model

import (
	"gorm.io/datatypes"
)

type UserChallenges struct {
	ID         uint           `gorm:"primarykey"`
	Address    string         `gorm:"column:address;type:varchar(44);index:,unique,composite:challenges_address_tokenId;comment:钱包地址" json:"address" form:"address"`
	TokenId    int64          `gorm:"column:token_id;index:,unique,composite:challenges_address_tokenId;" json:"tokenId"`
	Status     uint8          `gorm:"column:status;default:0;size:30;comment:0 进行中 1 等待验证 2 成功;" json:"status" form:"status"` // 0:进行中;1:等待验证;2:成功;
	Content    datatypes.JSON `gorm:"column:content" json:"content"`
	AddTs      int64          `gorm:"column:add_ts;autoCreateTime;comment:创建时间" json:"addTs"`
	Claimed    bool           `gorm:"column:claimed;default:false;comment:领取状态" json:"claimed"`
	UpdateTs   int64          `gorm:"column:update_ts;autoUpdateTime;comment:更新时间" json:"updateTs"`
	UserScore  int64          `gorm:"column:user_score;comment:分数" form:"user_score" json:"user_score"` // 分数
	ClaimTs    int64          `gorm:"column:claim_ts;comment:领取时间" json:"claimTs"`
	NFTAddress string         `gorm:"column:nft_address;comment:Solana NFT 地址" json:"nft_address" form:"nft_address"`
}
