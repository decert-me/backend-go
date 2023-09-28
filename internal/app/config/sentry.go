package config

type Sentry struct {
	Dsn              string  `mapstructure:"dsn" json:"dsn" yaml:"dsn"`
	TracesSampleRate float64 `mapstructure:"traces-sample-rate" json:"traces-sample-rate" yaml:"traces-sample-rate"`
	EnableTracing    bool    `mapstructure:"enable-tracing" json:"enable-tracing" yaml:"enable-tracing"`
}
