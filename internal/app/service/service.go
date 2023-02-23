package service

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/service/blockchain"
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// Service struct
type Service struct {
	c          *config.Config
	dao        *dao.Dao
	cron       *cron.Cron
	log        *zap.Logger
	blockchain *blockchain.BlockChain
}

// New init.
func New(c *config.Config, log *zap.Logger) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c, log),
		log: log,
	}
	s.blockchain = blockchain.New(c, s.dao, log)
	if s.c.Scheduler.Active {
		s.cron = cron.New()
		if _, err := s.cron.AddFunc(c.Scheduler.AirdropBadge, func() { s.blockchain.AirdropBadge() }); err != nil {
			s.log.Error("AirdropBadge cron init error", zap.Error(err))
		}
		s.cron.Start()
	}

	return
}

// Close Service.
func (s *Service) Close() {
	s.dao.Close()
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	err = s.dao.Ping(c)
	return
}
