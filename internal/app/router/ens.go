package router

import (
	v1 "backend-go/internal/app/api/v1"
	"github.com/gin-gonic/gin"
)

func InitEnsRouter(Router *gin.RouterGroup) {
	ensRouter := Router.Group("ens")
	{
		ensRouter.GET("/:q", v1.GetEnsRecords)
	}
}
