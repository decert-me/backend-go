package model

type ClaimBadgeTweet struct {
	ID          uint   `gorm:"primarykey"`
	Address     string `gorm:"column:address;type:char(42);UNIQUE;comment:钱包地址" json:"address" form:"address"`
	TokenId     string `gorm:"column:token_id;type:varchar" json:"tokenId"` // badgeNFT tokenId
	Url         string `gorm:"column:url;type:varchar" json:"url"`          // 推文链接地址
	AddTs       int64  `gorm:"autoCreateTime" json:"add_ts"`
	AirDroped   bool   `gorm:"" json:"airdroped"`
	AirdropTs   int64  `gorm:"" json:"airdrop_ts"`
	AirdropHash string `gorm:"type:varchar" json:"airdrop_hash"` // badgeNFT tokenId
}
