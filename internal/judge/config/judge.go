package config

type Judge struct {
	WorkPath  string `mapstructure:"work-path" json:"work-path" yaml:"work-path"`
	CachePath string `mapstructure:"cache-path" json:"cache-path" yaml:"cache-path"` // 缓存路径
}
