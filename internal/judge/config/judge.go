package config

type Judge struct {
	SandBoxService    string `mapstructure:"sandbox-service" json:"sandbox-service" yaml:"sandbox-service"`
	SolidityWorkPath  string `mapstructure:"solidity-work-path" json:"solidity-work-path" yaml:"solidity-work-path"`
	SolidityCachePath string `mapstructure:"solidity-cache-path" json:"solidity-cache-path" yaml:"solidity-cache-path"` // 缓存路径
	JavaScriptPath    string `mapstructure:"javascript-path" json:"javascript-path" yaml:"javascript-path"`             // JavaScript 路径
	TypeScriptPath    string `mapstructure:"typescript-path" json:"typescript-path" yaml:"typescript-path"`             // TypeScript 路径
	GolangPath        string `mapstructure:"golang-path" json:"golang-path" yaml:"golang-path"`                         // Golang 路径
	PythonPath        string `mapstructure:"python-path" json:"python-path" yaml:"python-path"`                         // Python 路径
}
