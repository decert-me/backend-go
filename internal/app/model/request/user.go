package request

type GetLoginMessageRequest struct {
	Address string `json:"address" form:"address"`
}

type AuthLoginSignRequest struct {
	Address   string `json:"address" form:"address"`
	Message   string `json:"message" form:"message"`
	Signature string `json:"signature" form:"signature"`
}
