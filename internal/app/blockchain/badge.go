package blockchain

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/global"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"go.uber.org/zap"
	"strings"
)

var badgeAbi abi.ABI

// initialize contract abi
func init() {
	contractAbi, err := abi.JSON(strings.NewReader(ABI.BadgeMetaData.ABI))
	if err != nil {
		global.LOG.Error("Failed to Load Abi", zap.Error(err))
		panic(err)
	}
	badgeAbi = contractAbi
}
