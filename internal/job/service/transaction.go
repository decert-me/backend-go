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
	"math"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var traversed atomic.Bool // 任务运行状态

type taskTx struct {
	task     *model.Transaction
	txMap    *sync.Map
	countMap *sync.Map
}

func (s *Service) StartTransaction() {
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			traversed.Store(false)
			log.Errorv("HandleTransaction error", zap.Any("error", err))
			time.Sleep(time.Second * 1)
			go s.StartTransaction()
		}
	}()

	if traversed.Load() {
		return
	}
	traversed.Store(true)

	var txMap sync.Map
	var countMap sync.Map
	// 循环
	for {
		fmt.Println("Transaction")
		// 超出扫描次数删除
		countMap.Range(func(key, value interface{}) bool {
			v, ok := value.(int)
			if ok && v > 10 {
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
		// 任务列表
		for _, trans := range transHashList {
			trans.Hash = strings.TrimSpace(trans.Hash)
			_, loaded := txMap.LoadOrStore(trans.Hash, trans)
			if loaded == false {
				s.TaskChain <- taskTx{task: &trans, txMap: &txMap, countMap: &countMap}
			}
		}
		time.Sleep(time.Second * 3)
	}
}

func (s *Service) consumeTransaction() {
	for {
		go s.handleTransactionReceipt(<-s.TaskChain)
	}
}

func (s *Service) handleTransactionReceipt(task taskTx) {
	provider := s.w.Next()
	hash := task.task.Hash
	// 错误处理
	defer func() {
		if err := recover(); err != nil {
			provider.OnInvokeFault()
			log.Errorv("HandleTransactionReceipt", zap.Any("err ", err))
			time.Sleep(time.Second * 3)
			//  控制尝试次数
			times, exist := task.countMap.LoadOrStore(hash, 1)
			if exist {
				v, ok := times.(int)
				if ok {
					task.countMap.Store(hash, v+1)
				}
			}
			s.handleTransactionReceipt(task)
		}
	}()
	client, err := ethclient.Dial(provider.Item)
	if err != nil {
		log.Errorv("dial error", zap.Any("error", err))
		panic("Error dial")
	}
	defer client.Close()

	var delay time.Duration
	for i := 0; i < s.c.BlockChain.Attempt; i++ {
		delay = time.Duration(math.Floor(float64(i)/50)*0.5 + 1)
		// 解析 Hash
		//fmt.Println(transHash)
		res, err := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
		provider.OnInvokeSuccess()
		// 待交易
		if err != nil {
			fmt.Println("wait for transaction", hash)
			time.Sleep(time.Second)
			continue
		}
		// 交易失败
		if res.Status == 0 {
			fmt.Println("fail for transaction", hash)
			s.handleTraverseStatus(hash, 2, "")
			task.txMap.Delete(hash)
			task.countMap.Delete(hash)
			return
		}
		// 交易成功
		if res.Status == 1 {
			fmt.Println("success for transaction", hash)
			if err = s.eventsParser(hash, res.Logs); err != nil {
				log.Errorv("EventsParser", zap.Any("err", err))
			}
			task.txMap.Delete(hash)
			task.countMap.Delete(hash)
			return
		}

		time.Sleep(delay * time.Second)
	}
	// 超出尝试次数
	s.handleTraverseStatus(hash, 3, "")
}

func (s *Service) eventsParser(hash string, Logs []*types.Log) (err error) {
	for _, vLog := range Logs {
		name, ok := s.contractEvent[vLog.Topics[0]]
		if !ok {
			continue
		}
		switch name {
		case "QuestCreated":
			if err = s.handleQuestCreated(hash, vLog); err != nil {
				s.handleTraverseStatus(hash, 5, err.Error())
				return err
			}
			return nil
		case "Claimed":
			if err = s.handleClaimed(hash, vLog); err != nil {
				s.handleTraverseStatus(hash, 5, err.Error())
				return err
			}
			return nil
		}
	}
	s.handleTraverseStatus(hash, 4, "")
	return nil
}

func (s *Service) handleTraverseStatus(hash string, status uint8, msg string) {
	err := s.dao.UpdateTransactionStatus(&model.Transaction{Hash: hash, Status: status, Msg: msg})
	if err != nil {
		log.Errorv("UpdateTransactionStatus error", zap.Error(err))
	}
}
