package v1

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func HandleProxy(c *gin.Context) {
	// 获取原始请求的URL和方法
	url := "http://127.0.0.1:8545"
	method := c.Request.Method

	// 创建新的请求
	req, err := http.NewRequest(method, url, c.Request.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// 设置请求头
	for k, v := range c.Request.Header {
		req.Header.Set(k, v[0])
	}

	// 发送请求并获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// 设置响应头
	for k, v := range resp.Header {
		c.Header(k, v[0])
	}

	// 设置响应状态码和响应体
	c.Status(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Writer.Write(body)
}
