package config

type Pgsql struct {
	Path          string `mapstructure:"path" json:"path" yaml:"path"`                               // 服务器地址
	Port          string `mapstructure:"port" json:"port" yaml:"port"`                               // 端口
	Config        string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname        string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username      string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password      string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	AutoMigrate   bool   `mapstructure:"auto-migrate" json:"auto-migrate" yaml:"auto-migrate"`       // 自动迁移
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 数据库表名前缀
	SlowThreshold int    `mapstructure:"slow-threshold" json:"slow-threshold" yaml:"slow-threshold"` // 慢查询阈值
	MaxIdleConns  int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns  int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode       string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap        bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

// Dsn 基于配置文件获取 dsn
func (p *Pgsql) Dsn() string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}

// LinkDsn 根据 dbname 生成 dsn
func (p *Pgsql) LinkDsn(dbname string) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}

func (m *Pgsql) GetLogMode() string {
	return m.LogMode
}