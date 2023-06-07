package env

import (
	"backend-go/sandbox/env/pool"
	"backend-go/sandbox/env/winc"
)

// NewBuilder build a environment builder
func NewBuilder(c Config) (pool.EnvBuilder, map[string]any, error) {
	b, err := winc.NewBuilder("")
	if err != nil {
		return nil, nil, err
	}
	c.Info("created winc builder")
	return b, map[string]any{}, nil
}
