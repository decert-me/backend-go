package initialize

import (
	ABI "backend-go/abi"
	ABIV2 "backend-go/abi/v2"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

// NewContractEvent 加载合约Event信息
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

// NewContractEventV2 加载合约Event信息
func NewContractEventV2() (contract map[common.Hash]string) {
	contract = make(map[common.Hash]string)
	ABIList := []string{ABIV2.BadgeV2MetaData.ABI, ABIV2.BadgeMinterV2MetaData.ABI, ABIV2.QuestV2MetaData.ABI, ABIV2.QuestMinterV2MetaData.ABI}
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
