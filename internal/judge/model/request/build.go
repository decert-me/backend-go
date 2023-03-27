package request

type BuildReq struct {
	Lang    string `json:"lang"`
	Code    string `json:"code"`
	Address string
}
