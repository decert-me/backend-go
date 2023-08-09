package config

type Share struct {
	VerifyKey string `mapstructure:"verify-key" json:"verify-key" yaml:"verify-key"` // 校验key
	Callback  string `mapstructure:"callback" json:"callback" yaml:"callback"`       // 回调
}
