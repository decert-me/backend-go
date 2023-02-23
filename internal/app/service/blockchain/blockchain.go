package blockchain

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/initialize"
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"sync/atomic"
)

type BlockChain struct {
	c             *config.Config
	dao           *dao.Dao
	client        *ethclient.Client
	traversed     atomic.Bool // 任务运行状态
	TaskChain     chan model.Transaction
	contractEvent map[common.Hash]string // 合约事件
}

func New(c *config.Config, dao *dao.Dao) (b *BlockChain) {
	b = &BlockChain{
		c:             c,
		dao:           dao,
		client:        new(ethclient.Client),
		TaskChain:     make(chan model.Transaction, 100),
		contractEvent: initialize.NewContractEvent(),
	}
	var err error
	b.client, err = ethclient.Dial(c.BlockChain.Provider)
	if err != nil {
		log.Error("ethclient Dial error", zap.Error(err))
	}
	b.StartTransaction()
	return
}
