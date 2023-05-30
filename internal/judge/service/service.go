package service

import (
	"backend-go/internal/judge/config"
	"backend-go/internal/judge/dao"
	"backend-go/pkg/log"
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// Service struct
type Service struct {
	c    *config.Config
	dao  *dao.Dao
	cron *cron.Cron
}

// New init.
func New(c *config.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	if s.c.Docker.ClearEnabled {
		s.cron = cron.New()
		if _, err := s.cron.AddFunc("*/1 * * * *", func() { s.ClearDocker() }); err != nil {
			log.Errorv("ClearDocker cron init error", zap.Error(err))
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
