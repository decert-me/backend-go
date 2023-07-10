package request

type TwitterAuthorizationReq struct {
	Callback string `json:"callback" form:"callback" query:"callback" binding:"required"`
}

type TwitterCallbackReq struct {
	Callback     string `json:"callback" binding:"required"`
	RequestToken string `json:"request_token" binding:"required"`
	Verifier     string `json:"verifier" binding:"required"`
}

type TwitterClaimReq struct {
	TokenId int64  `json:"tokenId"`
	Address string `json:"address"`
	Score   int64  `json:"score"  binding:"required"`
}
