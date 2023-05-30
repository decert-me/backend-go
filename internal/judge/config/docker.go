package config

type Docker struct {
	ClearEnabled bool `mapstructure:"clear-enabled" json:"clear-enabled" yaml:"clear-enabled"`
	ClearTime    int  `mapstructure:"clear-time" json:"clear-time" yaml:"clear-time"`
}
