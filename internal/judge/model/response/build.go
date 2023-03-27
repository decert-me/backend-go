package response

type BuildRes struct {
	Output          string `json:"output"`
	Gas             string `json:"gas"`
	ContractAddress string `json:"contract_address"`
	ABI             string `json:"abi"`
}
