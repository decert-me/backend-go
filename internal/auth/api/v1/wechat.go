package v1

import (
	"github.com/gin-gonic/gin"
)

// GetWechatQrcode 获取关注二维码
func GetWechatQrcode(c *gin.Context) {
	address := c.GetString("address")
	if address == "" {
		Fail(c)
		return
	}
	if data, err := srv.GetWechatQrcode(address); err != nil {
		Fail(c)
	} else {
		OkWithData(data, c)
	}
	return
}

// WechatService 微信服务器验证
func WechatService(c *gin.Context) {
	srv.WechatService(c)
}
