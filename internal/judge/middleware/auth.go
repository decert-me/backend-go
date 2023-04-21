package middleware

import (
	v1 "backend-go/internal/judge/api/v1"
	"backend-go/pkg/auth"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 鉴权头部信息： x-token
		token := c.Request.Header.Get("x-token")
		if token == "" {
			v1.FailWithDetailed(gin.H{"reload": true}, v1.GetMessage(c, "UnauthorizedAccess"), c)
			//FailWithDetailed(gin.H{"reload": true}, "授权已过期或非法访问1", c)
			c.Abort()
			return
		}
		// 解析token包含的信息
		claims, err := midAuth.ParseToken(token)
		if err != nil {
			if err == auth.TokenExpired {
				v1.FailWithDetailed(gin.H{"reload": true}, v1.GetMessage(c, "UnauthorizedAccess"), c)
				c.Abort()
				return
			}
			v1.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("address", claims.Address)
	}
}

func Addr() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 鉴权头部信息： x-token
		token := c.Request.Header.Get("x-token")
		if token != "" {
			// 解析token包含的信息
			claims, err := midAuth.ParseToken(token)
			if err != nil {
				fmt.Println(err)
				return
			}

			c.Set("address", claims.Address)
		}
	}
}
