package config

import (
	"backend-go/pkg/auth"
	"backend-go/pkg/log"
)

type Config struct {
	Log        *log.Config  `mapstructure:"log" json:"log" yaml:"log"`
	System     *System      `mapstructure:"system" json:"system" yaml:"system"`
	Pgsql      *Pgsql       `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	BlockChain *BlockChain  `mapstructure:"blockchain" json:"blockchain" yaml:"blockchain"`
	Contract   *Contract    `mapstructure:"contract" json:"contract" yaml:"contract"`
	Auth       *auth.Config `mapstructure:"auth" json:"auth" yaml:"auth"`
	Redis      *Redis       `mapstructure:"redis" json:"redis" yaml:"redis"`
	Quest      *Quest       `mapstructure:"quest" json:"quest" yaml:"quest"`
}
