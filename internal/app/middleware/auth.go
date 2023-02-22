package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		address := c.Request.Header.Get("address")
		c.Set("address", address)
	}
}
