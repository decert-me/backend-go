package config

import (
	"backend-go/pkg/auth"
	"backend-go/pkg/log"
)

type Config struct {
	Log        *log.Config        `mapstructure:"log" json:"log" yaml:"log"`
	System     *System            `mapstructure:"system" json:"system" yaml:"system"`
	Pgsql      *Pgsql             `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	BlockChain *BlockChain        `mapstructure:"blockchain" json:"blockchain" yaml:"blockchain"`
	Contract   *Contract          `mapstructure:"contract" json:"contract" yaml:"contract"`
	Auth       *auth.Config       `mapstructure:"auth" json:"auth" yaml:"auth"`
	Redis      *Redis             `mapstructure:"redis" json:"redis" yaml:"redis"`
	Quest      *Quest             `mapstructure:"quest" json:"quest" yaml:"quest"`
	Local      *Local             `mapstructure:"local" json:"local" yaml:"local"`
	IPFS       []IPFS             `mapstructure:"ipfs" json:"ipfs" yaml:"ipfs"`
	Judge      *Judge             `mapstructure:"judge" json:"judge" yaml:"judge"`
	Share      *Share             `mapstructure:"share" json:"share" yaml:"share"`
	Discord    *Discord           `mapstructure:"discord" json:"discord" yaml:"discord"`
	Sentry     *Sentry            `mapstructure:"sentry" json:"sentry" yaml:"sentry"`
	Social     *Social            `mapstructure:"social" json:"social" yaml:"social"`
	ZCloak     *ZCloak            `mapstructure:"zcloak" json:"zcloak" yaml:"zcloak"`
	NFT        *NFT               `mapstructure:"nft" json:"nft" yaml:"nft"`
	ContractV2 map[int64]Contract `mapstructure:"contract" json:"contract" yaml:"contract"`
}
