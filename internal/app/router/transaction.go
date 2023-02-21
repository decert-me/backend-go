package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitTransactionRouter(Router *gin.RouterGroup) {
	transactionRouter := Router.Group("transaction").Use(middleware.Auth())
	{
		transactionRouter.POST("submit", v1.HashSubmit)
	}
}
