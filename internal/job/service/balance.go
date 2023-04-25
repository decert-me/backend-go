package service

import (
	ABI "backend-go/abi"
	"backend-go/internal/app/utils"
	"backend-go/pkg/log"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"sync/atomic"
	"time"
)

var lock atomic.Bool

func (s *Service) BalanceRPC(chainID int) {
	if lock.Load() {
		return
	}
	lock.Store(true)
	defer lock.Store(false)
	providerList := s.c.BlockChain.Provider
	temp := make(map[int]string)
	for _, v := range providerList {
		if chainID != 0 && chainID != v.ChainID {
			temp[v.ChainID] = s.providerMap[v.ChainID]
			continue
		}
		if len(v.Url) == 0 {
			return
		}
		indexList := make([]int64, len(v.Url))
		for i, url := range v.Url {
			spent := s.rpcRequest(chainID, url)
			indexList[i] = spent
		}
		i, _ := utils.SliceMin[int64](indexList)
		temp[v.ChainID] = v.Url[i]
		log.Warn("RPC 切换: " + v.Name + " " + strconv.Itoa(i))
	}
	s.providerMap = temp
}

func (s *Service) rpcRequest(chainID int, url string) (spent int64) {
	max := int64(9999999999999)
	defer func() {
		if err := recover(); err != nil {
			spent = max
			return
		}
	}()
	startTime := time.Now()
	rpcClient, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	blockNumber, err := rpcClient.BlockNumber(ctx)
	if err != nil {
		return max
	}
	address := common.HexToAddress(s.c.Contract.MultiChain[chainID].Badge)
	instance, err := ABI.NewBadge(address, rpcClient)
	if err != nil {
		fmt.Println(err)
	}
	quest, err := instance.GetQuest(nil, big.NewInt(10000))
	if err != nil || quest.Title == "" || blockNumber == 0 {
		return max
	}
	return time.Since(startTime).Milliseconds()
}
