package initialize

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

// InitCommonDB 通用数据库
func InitCommonDB() {
	db := GormPgSql("")
	if db != nil {
		global.DB = db
		RegisterTables(db) // 初始化表
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.Users{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
