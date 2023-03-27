package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitIPFSRouter(Router *gin.RouterGroup) {
	ipfsRouterAuth := Router.Group("ipfs").Use(middleware.Auth())
	{
		ipfsRouterAuth.POST("uploadJson", v1.UploadJson)
		ipfsRouterAuth.POST("uploadFile", v1.UploadFile)
	}
}
