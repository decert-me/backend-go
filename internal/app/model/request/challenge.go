package request

type GetChallengeListRequest struct {
	PageInfo
	Address    string
	ReqAddress string
	Type       uint8 `form:"type"`
}

type SaveChallengeLogRequest struct {
	Address string `json:"-"`
	TokenId int64  `json:"token_id" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
}
