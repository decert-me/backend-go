package middleware

import (
	"backend-go/internal/app/config"
	"backend-go/pkg/auth"
)

var (
	midAuth *auth.Auth
)

func Init(c *config.Config) {
	midAuth = auth.New(c.Auth)
}
