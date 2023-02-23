package blockchain

import (
	ABI "backend-go/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

var badgeAbi abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.BadgeMetaData.ABI))
	if err != nil {
		panic(err)
	}
	badgeAbi = contractAbi
}
