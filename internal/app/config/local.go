package config

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 路径
	IPFS string `mapstructure:"ipfs" json:"ipfs" yaml:"ipfs"` // IPFS
}
