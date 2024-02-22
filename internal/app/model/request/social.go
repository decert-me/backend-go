package request

type GetEmailBindCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type EmailBindAddressRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}
