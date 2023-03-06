package config

type BlockChain struct {
	SignPrivateKey string `mapstructure:"sign-private-key" json:"sign-private-key" yaml:"sign-private-key"` // 签名私钥
	Signature      string `mapstructure:"signature" json:"signature" yaml:"signature"`                      // 签名信息
	SignatureExp   int    `mapstructure:"signature-exp" json:"signature-exp" yaml:"signature-exp"`          // 有效期
}
