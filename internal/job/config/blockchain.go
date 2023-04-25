package config

type Provider struct {
	Default bool     `mapstructure:"default" json:"default" yaml:"default"`    // 是否默认
	Name    string   `mapstructure:"name" json:"name" yaml:"name"`             // 名称
	ChainID int      `mapstructure:"chain-id" json:"chain-id" yaml:"chain-id"` // 链ID
	Url     []string `mapstructure:"url" json:"url" yaml:"url"`                // Provider Url
}

type BlockChain struct {
	SignPrivateKey    string              `mapstructure:"sign-private-key" json:"sign-private-key" yaml:"sign-private-key"`          // 签名私钥
	AirdropPrivateKey string              `mapstructure:"airdrop-private-key" json:"airdrop-private-key" yaml:"airdrop-private-key"` // Airdrop私钥
	ChainID           int64               `mapstructure:"chain-id" json:"chain-id" yaml:"chain-id"`                                  // Chain ID
	Attempt           int                 `mapstructure:"attempt" json:"attempt" yaml:"attempt"`                                     // 尝试次数
	Provider          map[string]Provider `mapstructure:"provider" json:"provider" yaml:"provider"`                                  // Provider
}
