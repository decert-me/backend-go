package v1

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

// claim Badge NFT
func PermitClaimBadge(c *gin.Context) {
	var req request.PermitClaimBadgeReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.PermitClaimBadge(address, req); err != nil {
		Fail(c)
	} else {
		OkWithData(list, c)
	}
}

func SubmitClaimTweet(c *gin.Context) {
	var req request.SubmitClaimTweetReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if err := srv.SubmitClaimTweet(address, req); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}
