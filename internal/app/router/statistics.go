package router

import (
	v1 "backend-go/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitStatisticsRouter(Router *gin.RouterGroup) {
	router := Router.Group("statistics")
	{
		router.GET("/addressChallengeCount/:address", v1.GetAddressChallengeCount) // 获取地址完成挑战/获得NFT的数量
	}
}
