package response

type CastCallRes struct {
	Data string `json:"data"`
}

type CastSend struct {
	Msg     string `json:"msg"`
	Status  string `json:"status"`
	GasUsed string `json:"gasUsed"`
}
