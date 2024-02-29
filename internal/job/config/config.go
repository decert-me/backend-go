package config

import (
	"backend-go/pkg/log"
)

type Config struct {
	Log        *log.Config          `mapstructure:"log" json:"log" yaml:"log"`
	Pgsql      *Pgsql               `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	BlockChain *BlockChain          `mapstructure:"blockchain" json:"blockchain" yaml:"blockchain"`
	Contract   *Contract            `mapstructure:"contract" json:"contract" yaml:"contract"`
	IPFS       *IPFS                `mapstructure:"ipfs" json:"ipfs" yaml:"ipfs"`
	ContractV2 map[int64]ContractV2 `mapstructure:"contract-v2" json:"contract-v2" yaml:"contract-v2"`
}
