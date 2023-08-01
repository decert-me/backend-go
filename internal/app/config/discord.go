package config

type Discord struct {
	Active         bool   `mapstructure:"active" json:"active" yaml:"active"`                            // 是否启用提醒功能
	Token          string `mapstructure:"token" json:"token" yaml:"token"`                               // 机器人Token
	SuccessChannel string `mapstructure:"success-channel" json:"success-channel" yaml:"success-channel"` // 成功提醒频道
	FailedChannel  string `mapstructure:"failed-channel" json:"failed-channel" yaml:"failed-channel"`    // 失败提醒频道
}
