package service

import (
	"backend-go/internal/job/config"
	"backend-go/internal/job/dao"
	"backend-go/internal/job/initialize"
	"backend-go/pkg/balancer"
	"backend-go/pkg/log"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// Service struct
type Service struct {
	c             *config.Config
	dao           *dao.Dao
	cron          *cron.Cron
	w             *balancer.SmoothRoundrobin
	TaskChain     chan taskTx
	contractEvent map[common.Hash]string // 合约事件
}

// New init.
func New(c *config.Config) (s *Service) {
	s = &Service{
		c:             c,
		dao:           dao.New(c),
		TaskChain:     make(chan taskTx, 100),
		contractEvent: initialize.NewContractEvent(),
	}
	s.w = initialize.InitProvider(c)
	go s.consumeTransaction() // 消费
	go s.StartTransaction()   // 任务
	if s.c.Scheduler.Active {
		s.cron = cron.New()
		if _, err := s.cron.AddFunc(c.Scheduler.AirdropBadge, func() { s.AirdropBadge() }); err != nil {
			log.Errorv("AirdropBadge cron init error", zap.Error(err))
		}
		s.cron.Start()
	}

	return
}

// Close Service.
func (s *Service) Close() {
	if s.cron != nil {
		<-s.cron.Stop().Done() // stop cron
	}
	s.dao.Close() // close db
	s = nil
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	err = s.dao.Ping(c)
	return
}
