package v1

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// claim Badge NFT
func PermitClaimBadge(c *gin.Context) {
	var req request.PermitClaimBadgeReq
	err := c.ShouldBindJSON(&req)
	fmt.Println(req)
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("param not valid", c)
		return
	}
	address := c.GetString("address")
	if list, err := service.PermitClaimBadge(address, req); err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))
		response.Fail(c)
	} else {
		response.OkWithData(list, c)
	}
}

func SubmitClaimTweet(c *gin.Context) {
	var req request.SubmitClaimTweetReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if err := service.SubmitClaimTweet(address, req); err != nil {
		global.LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}
