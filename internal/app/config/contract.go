package config

type Contract struct {
	Badge       string `mapstructure:"badge" json:"badge" yaml:"badge"`                      // 合约地址
	Quest       string `mapstructure:"quest" json:"quest" yaml:"quest"`                      // 合约地址
	QuestMinter string `mapstructure:"quest-minter" json:"quest-minter" yaml:"quest-minter"` // 合约地址
}
