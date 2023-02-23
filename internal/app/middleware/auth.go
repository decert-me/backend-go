package middleware

import (
	"backend-go/internal/app/model/response"
	"backend-go/pkg/auth"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 鉴权头部信息： x-token
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期或非法访问1", c)
			c.Abort()
			return
		}
		// 解析token包含的信息
		claims, err := midAuth.ParseToken(token)
		if err != nil {
			if err == auth.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("address", claims.Address)
	}
}
