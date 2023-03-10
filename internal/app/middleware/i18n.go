package middleware

import (
	"github.com/gin-gonic/gin"
)

func I18n() gin.HandlerFunc {
	return func(c *gin.Context) {
		// chain
		chain := c.Request.Header.Get("x-lang")
		if chain != "" {
			c.Set("lang", chain)
		} else {
			c.Set("lang", "zh-CN") // 默认语言
		}
	}
}
