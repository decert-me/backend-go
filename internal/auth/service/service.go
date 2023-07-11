package service

import (
	"backend-go/internal/auth/config"
	"backend-go/internal/auth/dao"
	"context"
)

// Service struct
type Service struct {
	c   *config.Config
	dao *dao.Dao
}

// New init.
func New(c *config.Config) (s *Service) {
	s = &Service{
		c:   c,
		dao: dao.New(c),
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
