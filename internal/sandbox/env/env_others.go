//go:build !windows && !linux && !darwin

package env

import (
	"errors"
	"runtime"

	"backend-go/sandbox/env/pool"
)

func NewBuilder(c Config) (pool.EnvBuilder, map[string]any, error) {
	return nil, nil, errors.New("environment is not support on this platform" + runtime.GOOS)
}
