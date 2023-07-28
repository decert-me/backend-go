package request

type GenerateShareRequest struct {
	Params string `json:"params"`
}

type ClickShareRequest struct {
	ShareCode string `json:"share_code"`
}

type AirdropCallbackRequest struct {
	Hash            string `json:"hash"`
	ContractAddress string `json:"contract_address"`
	Receiver        string `json:"receiver"`
	TokenId         string `json:"tokenId"`
	Params          string `json:"params"`
}
