package config

type System struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`             // 环境值
	Addr    int    `mapstructure:"addr" json:"addr" yaml:"addr"`          // 端口值
	I18n    string `mapstructure:"i18n" json:"i18n" yaml:"i18n"`          // 国际化文件
	Website string `mapstructure:"website" json:"website" yaml:"website"` // 网站文件路径
}
