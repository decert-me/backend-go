package service

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/service/blockchain"
	"backend-go/pkg/log"
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// Service struct
type Service struct {
	c          *config.Config
	dao        *dao.Dao
	cron       *cron.Cron
	blockchain *blockchain.BlockChain
}

// New init.
func New(c *config.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	s.blockchain = blockchain.New(c, s.dao)
	if s.c.Scheduler.Active {
		s.cron = cron.New()
		if _, err := s.cron.AddFunc(c.Scheduler.AirdropBadge, func() { s.blockchain.AirdropBadge() }); err != nil {
			log.Errorv("AirdropBadge cron init error", zap.Error(err))
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
