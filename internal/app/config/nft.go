package config

type NFT struct {
	API    string `mapstructure:"api" json:"api" yaml:"api"`             // API 地址
	APIKey string `mapstructure:"api-key" json:"api-key" yaml:"api-key"` // API Key
}
