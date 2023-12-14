package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitOpenQuestRouter(Router *gin.RouterGroup) {
	router := Router.Group("openQuest").Use(middleware.Auth())
	{
		router.POST("getUserOpenQuestList", v1.GetUserOpenQuestList) // 获取用户开放题列表
		router.POST("reviewOpenQuest", v1.ReviewOpenQuest)           // 审核开放题目
	}
}
