package config

import (
	"backend-go/pkg/auth"
	"backend-go/pkg/log"
)

type Config struct {
	Log     *log.Config  `mapstructure:"log" json:"log" yaml:"log"`
	System  *System      `mapstructure:"system" json:"system" yaml:"system"`
	Pgsql   *Pgsql       `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Auth    *auth.Config `mapstructure:"auth" json:"auth" yaml:"auth"`
	Foundry *Foundry     `mapstructure:"foundry" json:"foundry" yaml:"foundry"`
	Quest   *Quest       `mapstructure:"quest" json:"quest" yaml:"quest"`
}
