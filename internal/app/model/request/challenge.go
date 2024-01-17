package request

type Claimable struct {
	TokenId string `json:"token_id"`
	AddTs   int64  `json:"add_ts"`
}
type GetChallengeListRequest struct {
	PageInfo
	Address    string
	ReqAddress string
	Type       uint8  `form:"type"`
	Claimable  string `form:"claimable"`
	Language   string `form:"-"`
}

type SaveChallengeLogRequest struct {
	Address string `json:"-"`
	TokenId string `json:"token_id" binding:"required"`
	Answer  string `json:"answer" binding:"required"`
	IP      string `json:"-"`
	URI     string `json:"uri"`
}
