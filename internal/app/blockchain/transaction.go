package blockchain

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"sync/atomic"
	"time"
)

var Traversed atomic.Bool // 任务运行状态
var TransactionCh = make(chan model.Transaction, 100)

func StartTransaction() {
	go HandleTransaction()
	var transHashList []model.Transaction
	db := global.DB.Model(&model.Transaction{})
	if err := db.Where("status = 0").Find(&transHashList).Error; err != nil {
		return
	}
	for _, transHash := range transHashList {
		TransactionCh <- transHash
	}
}

func HandleTransaction() {
	if Traversed.Load() {
		return
	}
	Traversed.Store(true)
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			Traversed.Store(false)
			global.LOG.Error("HandleTransaction", zap.Any("err:", err))
			time.Sleep(time.Second * 3)
			go StartTransaction()
		}
	}()

	client, err := ethclient.Dial(global.CONFIG.BlockChain.Provider)
	if err != nil {
		panic("ethclient dial error")
	}
	for {
		go HandleTransactionReceipt(client, <-TransactionCh)
		time.Sleep(time.Second * 3)
	}
}

func HandleTransactionReceipt(client *ethclient.Client, transHash model.Transaction) {
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			global.LOG.Error("HandleTransactionReceipt", zap.Any("err ", err))
		}
	}()
	start := time.Now()
	defer func() {
		fmt.Println(start)
		fmt.Println(time.Now())
	}()
	for i := 0; i < 100; i++ {
		// 解析 Hash
		fmt.Println(transHash)
		res, err := client.TransactionReceipt(context.Background(), common.HexToHash(transHash.Hash))
		// 待交易
		if err != nil {
			fmt.Println(err)
			fmt.Println("wait for transaction")
			time.Sleep(time.Second)
			continue
		}
		// 交易失败
		if res.Status == 0 {
			fmt.Println("fail for transaction")
			if err = HandleTraverseStatus(transHash.Hash, 2, ""); err != nil {
				fmt.Println(err)
			}
		}
		// 交易成功
		if res.Status == 1 {
			fmt.Println("success for transaction")
			if err = EventsParser(transHash.Hash, res.Logs); err != nil {
				global.LOG.Error("EventsParser", zap.Any("err", err))
				return
			} else {
				return
			}
		}
		time.Sleep(time.Second)
	}
	HandleTraverseStatus(transHash.Hash, 3, "")
}

func HandleTraverseStatus(transHash string, status uint, msg string) error {
	if err := global.DB.Model(&model.Transaction{}).
		Where("hash = ?", transHash).
		Updates(map[string]interface{}{"status": status, "msg": msg}).Error; err != nil {
		return err
	}
	return nil
}

func EventsParser(hash string, Logs []*types.Log) (err error) {
	for _, vLog := range Logs {
		name, ok := global.ContractEvent[vLog.Topics[0]]
		if !ok {
			continue
		}
		switch name {
		case "QuestCreated":
			if err = handleQuestCreated(hash, vLog); err != nil {
				HandleTraverseStatus(hash, 5, err.Error())
				return err
			}
		case "Claimed":
			if err = handleClaimed(hash, vLog); err != nil {
				HandleTraverseStatus(hash, 5, err.Error())
				return err
			}
		}
	}
	HandleTraverseStatus(hash, 4, "")
	return nil
}
