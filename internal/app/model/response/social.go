package response

// BindingResponse 绑定提示
type BindingResponse struct {
	Success               bool   `json:"success"`                 // 是否绑定成功
	CurrentBindingAddress string `json:"current_binding_address"` // 当前绑定地址
}

type BindingResultResponse struct {
	Bound                 bool   `json:"bound"`
	CurrentBindingAddress string `json:"current_binding_address"` // 当前绑定地址
}
