package config

import "backend-go/pkg/auth"

type Auth struct {
	Auth    *auth.Config `mapstructure:"auth" json:"auth" yaml:"auth"`
	Wechat  Wechat       `mapstructure:"wechat" json:"wechat" yaml:"wechat"`
	Discord Discord      `mapstructure:"discord" json:"discord" yaml:"discord"`
	Github  Github       `mapstructure:"github" json:"github" yaml:"github"`
}

type Wechat struct {
	AppID          string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`
	AppSecret      string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"`
	Token          string `mapstructure:"token" json:"token" yaml:"token"`
	EncodingAESKey string `mapstructure:"encoding-aes-key" json:"encoding-aes-key" yaml:"encoding-aes-key"`
}

type Discord struct {
	ClientID     string `mapstructure:"client-id" json:"client-id" yaml:"client-id"`
	ClientSecret string `mapstructure:"client-secret" json:"client-secret" yaml:"client-secret"`
	RedirectURL  string `mapstructure:"redirect-url" json:"redirect-url" yaml:"redirect-url"`
}

type Github struct {
	ClientID     string `mapstructure:"client-id" json:"client-id" yaml:"client-id"`
	ClientSecret string `mapstructure:"client-secret" json:"client-secret" yaml:"client-secret"`
}
