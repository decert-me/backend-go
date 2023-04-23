package config

import (
	"backend-go/pkg/log"
)

type Config struct {
	Log        *log.Config `mapstructure:"log" json:"log" yaml:"log"`
	Pgsql      *Pgsql      `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	BlockChain *BlockChain `mapstructure:"blockchain" json:"blockchain" yaml:"blockchain"`
	Contract   *Contract   `mapstructure:"contract" json:"contract" yaml:"contract"`
	Twitter    *Twitter    `mapstructure:"twitter" json:"twitter" yaml:"twitter"`
	Scheduler  *Scheduler  `mapstructure:"scheduler" json:"scheduler" yaml:"scheduler"`
	IPFS       *IPFS       `mapstructure:"ipfs" json:"ipfs" yaml:"ipfs"`
}
