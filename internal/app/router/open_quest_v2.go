package router

import (
	v1 "backend-go/internal/app/api/v1"
	v2 "backend-go/internal/app/api/v2"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitOpenQuestV2Router(Router *gin.RouterGroup) {
	router := Router.Group("openQuest").Use(middleware.Auth())
	{
		router.POST("getUserOpenQuestList", v2.GetUserOpenQuestList)             // 获取用户开放题列表
		router.POST("getUserOpenQuestDetailList", v2.GetUserOpenQuestDetailList) // 审核开放题目
		router.POST("reviewOpenQuest", v1.ReviewOpenQuest)                       // 审核开放题目
	}
}
