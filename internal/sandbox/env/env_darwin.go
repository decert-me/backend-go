package env

import (
	"backend-go/sandbox/env/macsandbox"
	"backend-go/sandbox/env/pool"
)

var defaultRead = []string{
	"/",
}

var defaultWrite = []string{
	"/tmp",
	"/dev/null",
	"/var/tmp",
}

// NewBuilder build a environment builder
func NewBuilder(c Config) (pool.EnvBuilder, map[string]any, error) {
	b := macsandbox.NewBuilder("", defaultRead, defaultWrite, c.NetShare)
	c.Info("created mac sandbox at", "")
	return b, map[string]any{}, nil
}
