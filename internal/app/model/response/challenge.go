package response

import "backend-go/internal/app/model"

type GetChallengeListRes struct {
	model.Quest
	CompleteTs int64  `json:"complete_ts"`
	Claimed    bool   `json:"claimed"`
	NFTAddress string `json:"nft_address"`
}
