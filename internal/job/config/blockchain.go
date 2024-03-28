package config

type Provider struct {
	Url    string `mapstructure:"url" json:"url" yaml:"url"`          // Provider Url
	Weight int64  `mapstructure:"weight" json:"weight" yaml:"weight"` // 权重
}

type BlockChain struct {
	ChainID  int64      `mapstructure:"chain-id" json:"chain-id" yaml:"chain-id"` // Chain ID
	Attempt  int        `mapstructure:"attempt" json:"attempt" yaml:"attempt"`    // 尝试次数
	Provider []Provider `mapstructure:"provider" json:"provider" yaml:"provider"` // Provider
}
