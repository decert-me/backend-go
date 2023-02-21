package config

type Twitter struct {
	BearToken    string `mapstructure:"bear-token" json:"bear-token" yaml:"bear-token"`          // bear-token
	ClaimContent string `mapstructure:"claim-content" json:"claim-content" yaml:"claim-content"` // 分享推文
}
