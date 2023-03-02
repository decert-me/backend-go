package dao

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/initialize"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Dao dao.
type Dao struct {
	c     *config.Config
	db    *gorm.DB
	redis *redis.Client
}

// New init mysql db.
func New(c *config.Config) *Dao {
	return &Dao{
		c:     c,
		db:    initialize.NewPgSQL(c.Pgsql),
		redis: initialize.NewRedis(c.Redis),
	}
}

// Close close the resource.
func (d *Dao) Close() {
	d.redis.Close()
	db, _ := d.db.DB()
	db.Close()
}

// Ping dao ping
func (d *Dao) Ping(ctx context.Context) (err error) {
	db, _ := d.db.DB()
	if err = db.Ping(); err != nil {
		return
	}
	return d.redis.Ping(ctx).Err()
}

// DB returns the database
func (d *Dao) DB() *gorm.DB {
	return d.db
}
