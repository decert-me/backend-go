package v2

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

func SubmitClaimShareV2(c *gin.Context) {
	var req request.SubmitClaimShareV2Req
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if res, err := srv.SubmitClaimShareV2(address, req); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(res, c)
	}
}
