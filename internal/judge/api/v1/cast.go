package v1

import (
	"backend-go/internal/judge/model/request"
	"backend-go/pkg/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CastCall(c *gin.Context) {
	var req request.CastCallReq
	_ = c.ShouldBindJSON(&req)
	if res, err := srv.CastCall(common.HexToAddress("0").String(), req); err != nil {
		log.Errorv("err", zap.Error(err))
		Fail(c)
	} else {
		OkWithData(res, c)
	}
}

func CastSend(c *gin.Context) {
	var req request.CastSendReq
	_ = c.ShouldBindJSON(&req)
	if res, err := srv.CastSend(common.HexToAddress("0").String(), req); err != nil {
		log.Errorv("err", zap.Error(err))
		Fail(c)
	} else {
		OkWithData(res, c)
	}
}
