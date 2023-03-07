package config

type Scheduler struct {
	Active       bool   `mapstructure:"active" json:"active" yaml:"active"`                      // 是否启用定时任务
	AirdropBadge string `mapstructure:"airdrop-badge" json:"airdrop-badge" yaml:"airdrop-badge"` // 定时空投
}
