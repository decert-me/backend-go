package request

type TwitterCallbackReq struct {
	RequestToken string `json:"request_token" binding:"required"`
	Verifier     string `json:"verifier" binding:"required"`
}
