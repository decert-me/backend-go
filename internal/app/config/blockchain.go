package config

type BlockChain struct {
	PrivateKey   string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`       // 私钥
	Provider     string `mapstructure:"provider" json:"provider" yaml:"provider"`                // Provider
	ChainID      int64  `mapstructure:"chain-id" json:"chain-id" yaml:"chain-id"`                // Chain ID
	Signature    string `mapstructure:"signature" json:"signature" yaml:"signature"`             // 签名信息
	SignatureExp int    `mapstructure:"signature-exp" json:"signature-exp" yaml:"signature-exp"` // 有效期
}
