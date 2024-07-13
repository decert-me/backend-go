package request

type GetEmailBindCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type EmailBindAddressRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Code    string `json:"code" binding:"required"`
	Replace bool   `json:"replace"` // 是否替换
}

type UnbindRequest struct {
	Type string `json:"type" binding:"required"`
}
