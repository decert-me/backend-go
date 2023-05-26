package config

type Hardhat struct {
	Path          string `mapstructure:"path" json:"path" yaml:"path"`                                  // hardhat 项目临时路径
	SolcCachePath string `mapstructure:"solc-cache-path" json:"solc-cache-path" yaml:"solc-cache-path"` // solc 缓存路径
	ResourcePath  string `mapstructure:"resource-path" json:"resource-path" yaml:"resource-path"`       // hardhat 默认项目
}
