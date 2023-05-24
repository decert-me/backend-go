package config

type Twitter struct {
	ConsumerKey    string `mapstructure:"consumer-key" json:"consumer-key" yaml:"consumer-key"`
	ConsumerSecret string `mapstructure:"consumer-secret" json:"consumer-secret" yaml:"consumer-secret"`
	CallbackURL    string `mapstructure:"callback-url" json:"callback-url" yaml:"callback-url"`
	BearerToken    string `mapstructure:"bearer-token" json:"bearer-token" yaml:"bearer-token"`
}

type Github struct {
}

type Auth struct {
	Twitter Twitter `mapstructure:"twitter" json:"twitter" yaml:"twitter"`
	Github  Github  `mapstructure:"github" json:"github" yaml:"github"`
}
