package initialize

import (
	"backend-go/internal/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// GormPgSql 初始化 Postgresql 数据库
func GormPgSql(c *config.Pgsql) *gorm.DB {
	if c.Dbname == "" {
		panic("Postgres database name is required")
	}
	pgsqlConfig := postgres.Config{
		DSN:                  c.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	if db, err := gorm.Open(postgres.New(pgsqlConfig), _config(c)); err != nil {
		panic("database open failed")
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
		return db
	}
}

// config gorm 自定义配置
func _config(c *config.Pgsql) *gorm.Config {
	slowThreshold := 200
	if c.SlowThreshold > 200 {
		slowThreshold = c.SlowThreshold
	}
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, NamingStrategy: schema.NamingStrategy{
		TablePrefix:   c.Prefix,
		SingularTable: true,
	}}
	_default := logger.New(newWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: time.Duration(slowThreshold) * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch c.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func newWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}
