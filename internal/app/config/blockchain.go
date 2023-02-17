package config

type BlockChain struct {
	PrivateKey string `mapstructure:"private-key" json:"private-key" yaml:"private-key"` // 私钥
	Provider   string `mapstructure:"provider" json:"provider" yaml:"provider"`          // Provider
}
