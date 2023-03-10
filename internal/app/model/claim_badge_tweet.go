package model

type ClaimBadgeTweet struct {
	ID          uint   `gorm:"primarykey"`
	Address     string `gorm:"column:address;type:char(42);index:address_tokenId,UNIQUE;comment:钱包地址" json:"address" form:"address"`
	TokenId     int64  `gorm:"column:token_id;index:address_tokenId,UNIQUE" json:"tokenId"` // badgeNFT tokenId
	Score       int64  `gorm:"column:score" form:"score" json:"score"`                      // badgeNFT score
	Url         string `gorm:"column:url;type:varchar" json:"url"`                          // 推文链接地址
	TweetId     string `gorm:"column:tweet_id" json:"tweetId"`                              // 推文ID
	AddTs       int64  `gorm:"column:add_ts;autoCreateTime" json:"add_ts"`
	AirdropTs   int64  `gorm:"column:airdrop_ts" json:"airdrop_ts"`
	AirdropHash string `gorm:"column:airdrop_hash;type:varchar" json:"airdrop_hash"`
	Status      uint8  `gorm:"column:status;default:0" json:"status"` // 状态 0: 待空投 1: 成功空投 2: 出现错误
	Msg         string `gorm:"column:msg" json:"msg"`                 // 消息
}
