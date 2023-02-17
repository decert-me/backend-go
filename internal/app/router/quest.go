package router

import (
	v1 "backend-go/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitQuestRouter(Router *gin.RouterGroup) {
	questRouter := Router.Group("quests")
	{
		questRouter.GET("", v1.GetQuestList)
	}
}
