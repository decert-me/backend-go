package config

type MultiChain struct {
	Badge       string `mapstructure:"badge" json:"badge" yaml:"badge"`                      // 合约地址
	BadgeMinter string `mapstructure:"badge-minter" json:"badge-minter" yaml:"badge-minter"` // 合约地址
}
