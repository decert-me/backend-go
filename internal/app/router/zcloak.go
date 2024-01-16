package router

import (
	v1 "backend-go/internal/app/api/v1"
	"backend-go/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func InitZcloakRouter(Router *gin.RouterGroup) {
	router := Router.Group("zcloak").Use(middleware.Addr())
	{
		router.POST("saveSignAndDid", v1.SaveSignAndDid)              // 保存签名和DID
		router.GET("getAddressDid", v1.GetAddressDid)                 // 查询地址绑定的Did
		router.GET("getDidSignMessage", v1.GetDidSignMessage)         // 获取 DID 签名
		router.POST("generateCard", v1.GenerateCard)                  // 生成 card 信息
		router.GET("getKeyFileSignature", v1.GetKeyFileWithSignature) // 获取KeyFiles
		router.GET("getDidCardInfo", v1.GetDidCardInfo)               // 获取 DID card 信息
	}
}
