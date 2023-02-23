package config

type BlockChain struct {
	PrivateKey   string `mapstructure:"private-key" json:"private-key" yaml:"private-key"`       // 私钥
	Provider     string `mapstructure:"provider" json:"provider" yaml:"provider"`                // Provider
	Signature    string `mapstructure:"signature" json:"signature" yaml:"signature"`             // 签名信息
	SignatureExp int    `mapstructure:"scheduler-exp" json:"scheduler-exp" yaml:"scheduler-exp"` // 有效期
}
