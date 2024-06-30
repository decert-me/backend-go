package response

type GetAddressChallengeCountRes struct {
	UserID      int64  `gorm:"user_id" json:"user_id"`
	Address     string `gorm:"column:address;type:varchar(44);UNIQUE;comment:钱包地址" json:"address" form:"address"`
	Name        string `gorm:"name" json:"name"`
	Tags        string `json:"tags"`
	SuccessNum  int64  `json:"success_num"`   // 挑战成功数量
	FailNum     int64  `json:"fail_num"`      // 挑战失败数量
	ClaimNum    int64  `json:"claim_num"`     // 领取NFT数量
	NotClaimNum int64  `json:"not_claim_num"` // 未领取NFT数量
}
