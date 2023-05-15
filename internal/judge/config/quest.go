package config

type Quest struct {
	EncryptKey string `mapstructure:"encrypt-key" json:"encrypt-key" yaml:"encrypt-key"` // 加密密钥
}
