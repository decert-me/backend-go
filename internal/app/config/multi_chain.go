package config

type MultiChain struct {
	ChainID     int    `mapstructure:"chain-id" json:"chain-id" yaml:"chain-id"`             // 链ID
	Badge       string `mapstructure:"badge" json:"badge" yaml:"badge"`                      // 合约地址
	BadgeMinter string `mapstructure:"badge-minter" json:"badge-minter" yaml:"badge-minter"` // 合约地址
	Rpc         string `mapstructure:"rpc" json:"rpc" yaml:"rpc"`                            //  查询 RPC
}
