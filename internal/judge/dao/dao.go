package dao

import (
	"backend-go/internal/judge/config"
	"context"
	"time"
)

// Dao dao.
type Dao struct {
	c      *config.Config
	active map[string]time.Time
}

// New init mysql db.
func New(c *config.Config) *Dao {
	return &Dao{
		c:      c,
		active: make(map[string]time.Time),
	}
}

// Close close the resource.
func (d *Dao) Close() {
}

// Ping dao ping
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}
