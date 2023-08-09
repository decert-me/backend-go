package service

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/dao"
	"backend-go/internal/app/initialize"
	"backend-go/pkg/balancer"
	"context"
	"github.com/robfig/cron/v3"
)

// Service struct
type Service struct {
	c    *config.Config
	dao  *dao.Dao
	cron *cron.Cron
	W    *balancer.SmoothRoundrobin
}

// New init.
func New(c *config.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
	}
	s.W = initialize.InitJudge(c)
	return
}
func (s *Service) GetConfig() *config.Config {
	return s.c
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
