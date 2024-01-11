package service

import (
	"backend-go/internal/auth/config"
	"context"
)

// Service struct
type Service struct {
	c *config.Config
}

// New init.
func New(c *config.Config) (s *Service) {
	s = &Service{
		c: c,
	}

	return
}

// Close Service.
func (s *Service) Close() {
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return
}
