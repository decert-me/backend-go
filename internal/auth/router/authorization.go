package router

import (
	v1 "backend-go/internal/auth/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	authRouter := Router.Group("authorization")
	{
		authRouter.GET("twitter", v1.TwitterAuthorizationURL)
	}
}
