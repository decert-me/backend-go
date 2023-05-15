package initialize

import (
	"backend-go/internal/app/config"
	"backend-go/internal/app/model"
	"gorm.io/gorm"
)

// NewPgSQL new pgsql db
func NewPgSQL(c *config.Pgsql) *gorm.DB {
	db := GormPgSql(c)
	if c.AutoMigrate {
		RegisterTables(db) // 初始化表
	}
	return db
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Users{},
		model.ClaimBadgeTweet{},
		model.Quest{},
		model.UserChallenges{},
		model.Transaction{},
		model.Upload{},
		model.UserChallengeLog{},
		model.Ens{},
	)
	if err != nil {
		panic("register table failed")
	}
}
