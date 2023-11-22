package response

import "backend-go/internal/app/model"

type GetChallengeListRes struct {
	model.Quest
	CompleteTs            int64  `json:"complete_ts"`
	Claimed               bool   `json:"claimed"`
	NFTAddress            string `json:"nft_address"`
	OpenQuestReviewStatus uint8  `gorm:"column:open_quest_review_status" json:"open_quest_review_status" ` // // 评阅开放题状态 1 未审核 2 已审核
}
