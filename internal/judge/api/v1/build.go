package v1

import (
	"backend-go/internal/judge/model/request"
	"backend-go/pkg/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func BuildSolidity(c *gin.Context) {
	var req request.BuildReq
	_ = c.ShouldBindJSON(&req)
	if res, err := srv.BuildSolidity("", req); err != nil {
		log.Errorv("err", zap.Error(err))
		Fail(c)
	} else {
		OkWithData(res, c)
	}
}
