package config

type System struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`    // 环境值
	Addr int    `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
}
