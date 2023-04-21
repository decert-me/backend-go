package response

type BuildRes struct {
	Output          string `json:"output"`
	Gas             string `json:"gas"`
	ContractAddress string `json:"contract_address"`
	ABI             string `json:"abi"`
	Status          uint8  `json:"status"` // 0 成功 1 异常
}
