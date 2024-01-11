package request

type HashSubmitRequest struct {
	Hash    string `json:"hash"`
	Params  string `json:"params"`
	ChainID int64  `json:"chain_id"`
}
