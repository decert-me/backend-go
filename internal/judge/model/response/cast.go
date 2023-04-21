package response

type CastCallRes struct {
	Data   string `json:"data"`
	Msg    string `json:"msg"`
	Status uint8  `json:"status"` // 0 成功 1 异常
}

type CastSend struct {
	Msg     string `json:"msg"`
	Status  string `json:"status"`
	GasUsed string `json:"gasUsed"`
}
