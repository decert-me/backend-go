package response

import "backend-go/internal/app/model"

type GetChallengeListRes struct {
	model.Quest
	Claimable             bool   `json:"claimable"`
	CompleteTs            int64  `gorm:"-" json:"complete_ts"`
	Claimed               bool   `gorm:"claimed" json:"claimed"`
	NFTAddress            string `gorm:"nft_address" json:"nft_address"`
	OpenQuestReviewStatus uint8  `gorm:"open_quest_review_status" json:"open_quest_review_status" `                                  // // 评阅开放题状态 1 未审核 2 已审核
	IsOpenQuest           bool   `gorm:"column:is_open_quest;default:false;comment:是否开放题" json:"is_open_quest" form:"is_open_quest"` // 是否开放题
	ClaimStatus           uint8  `gorm:"claim_status" json:"claim_status"`                                                           // 0 未领取 1 NFT 2 zcloak 3 两者
	BadgeTokenID          string `gorm:"column:badge_token_id" json:"badge_token_id"`
	BadgeChainID          int64  `gorm:"column:badge_chain_id" json:"badge_chain_id"`
}
