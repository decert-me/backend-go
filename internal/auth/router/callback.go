package router

import (
	"github.com/gin-gonic/gin"
)

func InitCallbackRouter(Router *gin.RouterGroup) {
	callbackRouter := Router.Group("callback")
	{
		_ = callbackRouter
	}
}
