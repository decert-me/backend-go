package request

type TryRunReq struct {
	Lang      string `json:"lang"`
	TokenID   uint64 `json:"token_id"`
	Code      string `json:"code"`
	DataInput string `json:"data_input"`
}
