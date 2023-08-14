package router

import (
	v1 "backend-go/internal/auth/api/v1"
	"github.com/gin-gonic/gin"
)

func InitTwitterRouter(Router *gin.RouterGroup) {
	twitterRouter := Router.Group("twitter")
	{
		twitterRouter.POST("tweet", v1.TwitterUserTweet)
	}
}
