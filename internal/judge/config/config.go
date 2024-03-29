package config

import (
	"backend-go/pkg/auth"
	"backend-go/pkg/log"
)

type Config struct {
	Log    *log.Config  `mapstructure:"log" json:"log" yaml:"log"`
	System *System      `mapstructure:"system" json:"system" yaml:"system"`
	Auth   *auth.Config `mapstructure:"auth" json:"auth" yaml:"auth"`
	Judge  *Judge       `mapstructure:"judge" json:"judge" yaml:"judge"`
	Quest  *Quest       `mapstructure:"quest" json:"quest" yaml:"quest"`
	Docker *Docker      `mapstructure:"docker" json:"docker" yaml:"docker"`
}
