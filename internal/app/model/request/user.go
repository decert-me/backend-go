package request

type GetLoginMessageRequest struct {
	Address string `json:"address" form:"address"`
}

type AuthLoginSignRequest struct {
	Address   string `json:"address" form:"address" binding:"required"`
	Message   string `json:"message" form:"message" binding:"required"`
	Signature string `json:"signature" form:"signature" binding:"required"`
}
