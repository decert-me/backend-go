package initialize

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/blockchain"
	"backend-go/internal/app/global"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// InitContract 加载合约信息
func InitContract() {
	initContractEvent()
	go blockchain.StartTransaction()
}

// initContractEvent 加载合约Event信息
func initContractEvent() {
	global.ContractEvent = make(map[common.Hash]string)
	ABIList := []string{ABI.BadgeMetaData.ABI, ABI.QuestMetaData.ABI, ABI.QuestMinterMetaData.ABI}
	for _, abiStr := range ABIList {
		contractAbi, err := abi.JSON(strings.NewReader(abiStr))
		if err != nil {
			panic(err)
		}
		for _, v := range contractAbi.Events {
			global.ContractEvent[v.ID] = v.Name
		}
	}
}
