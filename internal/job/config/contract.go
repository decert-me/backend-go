package config

type ContractAddress struct {
	Badge       string `mapstructure:"badge" json:"badge" yaml:"badge"`                      // 合约地址
	Quest       string `mapstructure:"quest" json:"quest" yaml:"quest"`                      // 合约地址
	QuestMinter string `mapstructure:"quest-minter" json:"quest-minter" yaml:"quest-minter"` // 合约地址
	BadgeMinter string `mapstructure:"badge-minter" json:"badge-minter" yaml:"badge-minter"` // 合约地址
}

type Contract struct {
	V1 ContractAddress `mapstructure:"v1" json:"v1" yaml:"v1"`
	V2 ContractAddress `mapstructure:"v2" json:"v2" yaml:"v2"`
}
