package initialize

import (
	ABI "backend-go/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// InitContract 加载合约信息
func InitContract() {
	//go blockchain.StartTransaction()
}

// initContractEvent 加载合约Event信息
func NewContractEvent() (contract map[common.Hash]string) {
	contract = make(map[common.Hash]string)
	ABIList := []string{ABI.BadgeMetaData.ABI, ABI.QuestMetaData.ABI, ABI.QuestMinterMetaData.ABI}
	for _, abiStr := range ABIList {
		contractAbi, err := abi.JSON(strings.NewReader(abiStr))
		if err != nil {
			panic(err)
		}
		for _, v := range contractAbi.Events {
			contract[v.ID] = v.Name
		}
	}
	return
}
