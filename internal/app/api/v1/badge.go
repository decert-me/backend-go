package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

// claim Badge NFT
func PermitClaimBadge(c *gin.Context) {
	var req request.PermitClaimBadgeReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("param not valid", c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.PermitClaimBadge(address, req); err != nil {
		response.Fail(c)
	} else {
		response.OkWithData(list, c)
	}
}

func SubmitClaimTweet(c *gin.Context) {
	var req request.SubmitClaimTweetReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if err := srv.SubmitClaimTweet(address, req); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}
}
