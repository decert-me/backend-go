package request

type TwitterCallbackReq struct {
	RequestToken string `json:"request_token" binding:"required"`
	Verifier     string `json:"verifier" binding:"required"`
}

type TwitterClaimReq struct {
	TokenId int64  `json:"tokenId"`
	Address string `json:"address"`
	Score   int64  `json:"score"  binding:"required"`
}
