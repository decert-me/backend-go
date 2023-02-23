package config

type Config struct {
	Zap        *Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
	System     *System     `mapstructure:"system" json:"system" yaml:"system"`
	Pgsql      *Pgsql      `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	BlockChain *BlockChain `mapstructure:"blockchain" json:"blockchain" yaml:"blockchain"`
	Contract   *Contract   `mapstructure:"contract" json:"contract" yaml:"contract"`
	Twitter    *Twitter    `mapstructure:"twitter" json:"twitter" yaml:"twitter"`
	JWT        *JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis      *Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	Scheduler  *Scheduler  `mapstructure:"scheduler" json:"scheduler" yaml:"scheduler"`
}
