package config

type Project struct {
	APIKey      string `mapstructure:"api-key" json:"api-key" yaml:"api-key"`                // API Key
	CallBackURL string `mapstructure:"callback-url" json:"callback-url" yaml:"callback-url"` // 回调地址
	//DiscordCallBackURL string `mapstructure:"discord-callback-url" json:"discord-callback-url" yaml:"discord-callback-url"` // 回调地址
}
