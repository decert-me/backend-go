package service

import (
	"backend-go/internal/auth/config"
	"context"
	"github.com/allegro/bigcache/v3"
	"time"
)

// Service struct
type Service struct {
	c     *config.Config
	cache *bigcache.BigCache
}

// New init.
func New(c *config.Config) (s *Service) {
	s = &Service{
		c: c,
	}
	s.cache, _ = bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	return
}

// Close Service.
func (s *Service) Close() {
}

// Ping check server ok.
func (s *Service) Ping(c context.Context) (err error) {
	return
}
