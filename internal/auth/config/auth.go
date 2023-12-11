package config

import "backend-go/pkg/auth"

type Twitter struct {
	ConsumerKey    string `mapstructure:"consumer-key" json:"consumer-key" yaml:"consumer-key"`
	ConsumerSecret string `mapstructure:"consumer-secret" json:"consumer-secret" yaml:"consumer-secret"`
	CallbackURL    string `mapstructure:"callback-url" json:"callback-url" yaml:"callback-url"`
}

type Github struct {
}

type Auth struct {
	Auth    *auth.Config `mapstructure:"auth" json:"auth" yaml:"auth"`
	Twitter Twitter      `mapstructure:"twitter" json:"twitter" yaml:"twitter"`
	Github  Github       `mapstructure:"github" json:"github" yaml:"github"`
	Wechat  Wechat       `mapstructure:"wechat" json:"wechat" yaml:"wechat"`
	Discord Discord      `mapstructure:"discord" json:"discord" yaml:"discord"`
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
