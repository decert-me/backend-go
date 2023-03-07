package dao

import (
	"backend-go/internal/job/config"
	"backend-go/internal/job/initialize"
	"context"
	"gorm.io/gorm"
)

// Dao dao.
type Dao struct {
	c  *config.Config
	db *gorm.DB
}

// New init mysql db.
func New(c *config.Config) *Dao {
	return &Dao{
		c:  c,
		db: initialize.NewPgSQL(c.Pgsql),
	}
}

// Close close the resource.
func (d *Dao) Close() {
	db, _ := d.db.DB()
	db.Close()
}

// Ping dao ping
func (d *Dao) Ping(ctx context.Context) (err error) {
	db, _ := d.db.DB()
	if err = db.Ping(); err != nil {
		return
	}
	return
}

// DB returns the database
func (d *Dao) DB() *gorm.DB {
	return d.db
}
