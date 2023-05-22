package config

type Github struct {
	ConsumerKey    string `mapstructure:"consumer-key" json:"consumer-key" yaml:"consumer-key"`
	ConsumerSecret string `mapstructure:"consumer-secret" json:"consumer-secret" yaml:"consumer-secret"`
	CallbackURL    string `mapstructure:"callback-url" json:"callback-url" yaml:"callback-url"`
}

type Auth struct {
	Github Github `mapstructure:"github" json:"github" yaml:"github"`
}
