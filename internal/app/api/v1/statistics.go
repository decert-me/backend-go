package v1

import (
	"github.com/gin-gonic/gin"
)

// GetAddressChallengeCount 获取地址完成挑战/获得NFT的数量
func GetAddressChallengeCount(c *gin.Context) {
	address := c.Param("address")
	if count, err := srv.GetAddressChallengeCount(address); err != nil {
		FailWithMessage("获取失败："+err.Error(), c)
	} else {
		OkWithData(count, c)
	}
}
