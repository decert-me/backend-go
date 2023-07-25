package config

type Share struct {
	ShareCallback string `mapstructure:"share-callback" json:"share-callback" yaml:"share-callback"` // 分享码回调
	ClickCallback string `mapstructure:"click-callback" json:"click-callback" yaml:"click-callback"` // 点击回调
}
