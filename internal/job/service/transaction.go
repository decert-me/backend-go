package service

import (
	"backend-go/internal/app/model"
	"backend-go/pkg/log"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

var Traversed sync.Map // 任务运行状态

type taskTx struct {
	task     *model.Transaction
	txMap    *sync.Map
	countMap *sync.Map
}

func (s *Service) HandleTransaction(chainID int) {
	_, ok := Traversed.Load(chainID)
	if ok {
		return
	}
	Traversed.Store(chainID, true)
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			Traversed.Delete(chainID)
			log.Errorv("HandleTransaction error", zap.Any("error", err))
			time.Sleep(time.Second * 1)
			go s.HandleTransaction(chainID)
		}
	}()

	client, err := ethclient.Dial(s.providerMap[chainID])
	if err != nil {
		panic("Error dial")
	}

	var txMap sync.Map
	var countMap sync.Map
	// 循环
	for {
		fmt.Println("Transaction")
		// 超出扫描次数删除
		countMap.Range(func(key, value interface{}) bool {
			v, ok := value.(int)
			if ok && v > 100 {
				fmt.Println("超出扫描次数删除")
				s.handleTraverseStatus(key.(string), 3, "")
				countMap.Delete(key)
			}
			return true
		})
		// 获取需要扫描的数据
		transHashList, err := s.dao.QueryWaitTransaction()
		if err != nil {
			log.Errorv("QueryWaitTransaction error", zap.Any("error", err))
			time.Sleep(time.Second * 3)
			continue
		}
		var haveBool bool // 是否空map
		txMap.Range(func(key, value interface{}) bool {
			haveBool = true
			return false
		})
		// 无任务
		if len(transHashList) == 0 && !haveBool {
			Traversed.Delete(chainID)
			return
		}
		// 任务列表
		for _, trans := range transHashList {
			trans.Hash = strings.TrimSpace(trans.Hash)
			_, loaded := txMap.LoadOrStore(trans.Hash, trans)
			if loaded == false {
				go s.handleTransactionReceipt(client, chainID, &txMap, &countMap, trans.Hash)
				//s.TaskChain <- taskTx{task: &trans, txMap: &txMap, countMap: &countMap}
			}
		}
		time.Sleep(time.Second * 3)
	}
}

func (s *Service) handleTransactionReceipt(client *ethclient.Client, chainID int, txMap *sync.Map, countMap *sync.Map, hash string) {
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			txMap.Delete(hash)
			log.Error("HandleTransactionReceipt致命错误", zap.Any("err ", err))
		}
	}()
	// 是否在处理列表
	transHashAny, ok := txMap.Load(hash)
	if !ok {
		return
	}
	transHash, ok := transHashAny.(model.Transaction)
	if !ok {
		log.Error("HandleTransactionReceipt Reflect Error")
		return
	}
	fmt.Println(hash)
	// 解析交易Hash
	res, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	// 待交易
	if err != nil {
		fmt.Println("待交易", err)
		txMap.Delete(hash)
		times, exist := countMap.LoadOrStore(hash, 1)
		if exist {
			v, ok := times.(int)
			if ok {
				countMap.Store(hash, v+1)
			}
		}
		return
	}
	// 交易失败
	if res.Status == 0 {
		fmt.Println("交易失败")
		txMap.Delete(hash)
		countMap.Delete(hash)
		s.handleTraverseStatus(transHash.Hash, 3, "")
		return
	}
	// 交易成功
	if res.Status == 1 {
		fmt.Println("交易成功")
		if err := s.eventsParser(transHash.Hash, res.Logs); err != nil {
			fmt.Println(err)
			txMap.Delete(hash)
		} else {
			fmt.Println("交易成功--删除")
			txMap.Delete(hash)
			countMap.Delete(hash)
		}
	}
}

func (s *Service) eventsParser(hash string, Logs []*types.Log) (err error) {
	for _, vLog := range Logs {
		name, ok := s.contractEvent[vLog.Topics[0]]
		if !ok {
			continue
		}
		fmt.Println(name)
		switch name {
		case "QuestCreated":
			if err := s.handleQuestCreated(hash, vLog); err != nil {
				s.handleTraverseStatus(hash, 5, err.Error())
				continue
			}
			return nil
		case "Claimed":
			if err := s.handleClaimed(hash, vLog); err != nil {
				s.handleTraverseStatus(hash, 5, err.Error())
				continue
			}
			return nil
		case "URI":
			if err := s.handleURI(hash, vLog); err != nil {
				s.handleTraverseStatus(hash, 5, err.Error())
				continue
			}
			return nil
		}

	}
	return nil
}

func (s *Service) handleTraverseStatus(hash string, status uint8, msg string) {
	err := s.dao.UpdateTransactionStatus(&model.Transaction{Hash: hash, Status: status, Msg: msg})
	if err != nil {
		log.Errorv("UpdateTransactionStatus error", zap.Error(err))
	}
}
