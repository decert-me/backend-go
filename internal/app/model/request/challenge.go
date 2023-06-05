package request

type Claimable struct {
	TokenId int64 `json:"token_id"`
	AddTs   int64 `json:"add_ts"`
}
type GetChallengeListRequest struct {
	PageInfo
	Address    string
	ReqAddress string
	Type       uint8  `form:"type"`
	Claimable  string `form:"claimable"`
}

type SaveChallengeLogRequest struct {
	Address string `json:"-"`
	TokenId int64  `json:"token_id" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
}
