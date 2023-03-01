package blockchain

import (
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"math"
	"time"
)

func (b *BlockChain) StartTransaction() {
	go b.handleTransaction()
	var transHashList []model.Transaction
	transHashList, err := b.dao.QueryWaitTransaction()
	if err != nil {
		log.Errorv("Error querying transaction", zap.Error(err))
	}
	for _, transHash := range transHashList {
		b.TaskChain <- transHash
	}
}

func (b *BlockChain) handleTransaction() {
	if b.traversed.Load() {
		return
	}
	b.traversed.Store(true)
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			b.traversed.Store(false)
			log.Errorv("HandleTransaction", zap.Any("err:", err))
			time.Sleep(time.Second * 3)
			go b.StartTransaction()
		}
	}()
	for {
		go b.handleTransactionReceipt(b.client, <-b.TaskChain)
		time.Sleep(time.Second * 3)
	}
}

func (b *BlockChain) handleTransactionReceipt(client *ethclient.Client, transHash model.Transaction) {
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			log.Errorv("HandleTransactionReceipt", zap.Any("err ", err))
		}
	}()
	var delay time.Duration
	for i := 0; i < 300; i++ {
		delay = time.Duration(math.Floor(float64(i)/50)*0.5 + 1)
		// 解析 Hash
		//fmt.Println(transHash)
		res, err := client.TransactionReceipt(context.Background(), common.HexToHash(transHash.Hash))
		// 待交易
		if err != nil {
			fmt.Println("wait for transaction", transHash.Hash)
			time.Sleep(time.Second)
			continue
		}
		// 交易失败
		if res.Status == 0 {
			fmt.Println("fail for transaction", transHash.Hash)
			b.handleTraverseStatus(transHash.Hash, 2, "")
			return
		}
		// 交易成功
		if res.Status == 1 {
			fmt.Println("success for transaction", transHash.Hash)
			if err = b.eventsParser(transHash.Hash, res.Logs); err != nil {
				log.Errorv("EventsParser", zap.Any("err", err))
				return
			} else {
				return
			}
		}
		time.Sleep(delay * time.Second)
	}
	// 超出尝试次数
	transHash.Status = 3

}

func (b *BlockChain) eventsParser(hash string, Logs []*types.Log) (err error) {
	for _, vLog := range Logs {
		name, ok := b.contractEvent[vLog.Topics[0]]
		if !ok {
			continue
		}
		switch name {
		case "QuestCreated":
			if err = b.handleQuestCreated(hash, vLog); err != nil {
				b.handleTraverseStatus(hash, 5, err.Error())
				return err
			}
			return nil
		case "Claimed":
			if err = b.handleClaimed(hash, vLog); err != nil {
				b.handleTraverseStatus(hash, 5, err.Error())
				return err
			}
			return nil
		}
	}
	b.handleTraverseStatus(hash, 4, "")
	return nil
}

func (b *BlockChain) handleTraverseStatus(hash string, status uint8, msg string) {
	err := b.dao.UpdateTransactionStatus(&model.Transaction{Hash: hash, Status: status, Msg: msg})
	if err != nil {
		log.Errorv("UpdateTransactionStatus error", zap.Error(err))
	}
}
