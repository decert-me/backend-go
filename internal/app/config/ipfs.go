package config

type IPFS struct {
	API       string `mapstructure:"api" json:"api" yaml:"api"`
	UploadAPI string `mapstructure:"upload-api" json:"upload-api" yaml:"upload-api"`
}
