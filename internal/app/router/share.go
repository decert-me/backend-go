package router

import (
	v1 "backend-go/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitShareRouter(Router *gin.RouterGroup) {
	shareRouter := Router.Group("share")
	{
		shareRouter.POST("click", v1.ClickShare)
		shareRouter.POST("callback", v1.AirdropCallback)
	}
}
