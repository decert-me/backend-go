package v1

import (
	"backend-go/internal/auth/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ERROR   = 7
	SUCCESS = 0
)

var (
	srv  *service.Service
	i18n map[string]map[string]string
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Init init
func Init(s *service.Service, i map[string]map[string]string) {
	srv = s
	i18n = i
}

func Ping(c *gin.Context) {
	if err := srv.Ping(c); err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, "ok")
}

// GetMessage get i18n message
func GetMessage(c *gin.Context, key string) string {
	lang := c.GetString("lang")
	languageMessages, ok := i18n[lang]
	if !ok {
		return key
	}

	message, ok := languageMessages[key]
	if !ok {
		return key
	}

	return message
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, GetMessage(c, "OperationSuccess"), c)
}

func OkWithRaw(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, GetMessage(c, "OperationSuccess"), c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, GetMessage(c, "OperationFailed"), c)
}
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
