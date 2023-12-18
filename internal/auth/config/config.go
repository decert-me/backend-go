package config

import (
	"backend-go/pkg/log"
)

type Config struct {
	Log     *log.Config        `mapstructure:"log" json:"log" yaml:"log"`
	System  *System            `mapstructure:"system" json:"system" yaml:"system"`
	Pgsql   *Pgsql             `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Auth    *Auth              `mapstructure:"auth" json:"auth" yaml:"auth"`
	Project map[string]Project `mapstructure:"project" json:"project" yaml:"project"`
}
