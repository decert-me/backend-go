package config

type Contract struct {
	Quest       string                `mapstructure:"quest" json:"quest" yaml:"quest"`                      // 合约地址
	QuestMinter string                `mapstructure:"quest-minter" json:"quest-minter" yaml:"quest-minter"` // 合约地址
	Badge       string                `mapstructure:"badge" json:"badge" yaml:"badge"`                      // 合约地址
	BadgeMinter string                `mapstructure:"badge-minter" json:"badge-minter" yaml:"badge-minter"` // 合约地址
	MultiChain  map[string]MultiChain `mapstructure:"multi-chain" json:"multi-chain" yaml:"multi-chain"`    // 多链合约地址
}
