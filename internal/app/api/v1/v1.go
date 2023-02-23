package v1

import (
	"backend-go/internal/app/service"
)

var (
	srv *service.Service
)

// Init init
func Init(s *service.Service) {
	srv = s
}
