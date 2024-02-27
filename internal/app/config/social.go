package config

type Social struct {
	Wechat  Wechat      `mapstructure:"wechat" json:"wechat" yaml:"wechat"`
	Discord DiscordBind `mapstructure:"discord" json:"discord" yaml:"discord"`
	Github  GithubBind  `mapstructure:"github" json:"github" yaml:"github"`
}

type Wechat struct {
	APIKey  string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`    // API Key
	CallURL string `mapstructure:"call-url" json:"call-url" yaml:"call-url"` // 调用地址
}

type DiscordBind struct {
	APIKey  string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`    // API Key
	CallURL string `mapstructure:"call-url" json:"call-url" yaml:"call-url"` // 调用地址
}

type GithubBind struct {
	APIKey  string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`    // API Key
	CallURL string `mapstructure:"call-url" json:"call-url" yaml:"call-url"` // 调用地址
}
