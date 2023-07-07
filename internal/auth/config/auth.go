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
}
