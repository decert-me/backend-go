package request

type CastCallReq struct {
	To       string `json:"to"`
	Method   string `json:"method"`
	Data     string `json:"data"`
	CallData string `json:"calldata"`
}

type CastSendReq struct {
	To       string `json:"to"`
	Method   string `json:"method"`
	Data     string `json:"data"`
	CallData string `json:"calldata"`
}
