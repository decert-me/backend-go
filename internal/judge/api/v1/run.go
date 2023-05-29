package v1

import (
	"backend-go/internal/judge/model/request"
	"backend-go/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TryRun(c *gin.Context) {
	var req request.TryRunReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if res, err := srv.TryRun(address, req); err != nil {
		log.Errorv("err", zap.Error(err))
		Fail(c)
	} else {
		OkWithData(res, c)
	}
}

func TryTestRun(c *gin.Context) {
	var req request.TryTestRunReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if res, err := srv.TryTestRun(address, req); err != nil {
		log.Errorv("err", zap.Error(err))
		Fail(c)
	} else {
		OkWithData(res, c)
	}
}
