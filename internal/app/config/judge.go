package config

type JudgeApi struct {
	Url    string `mapstructure:"url" json:"url" yaml:"url"`          // Provider Url
	Weight int64  `mapstructure:"weight" json:"weight" yaml:"weight"` // 权重
}

type Judge struct {
	JudgeApi []JudgeApi `mapstructure:"judge-api" json:"judge-api" yaml:"judge-api"`
}
