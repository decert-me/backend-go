package v1

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

// PermitClaimBadge claim Badge NFT
func PermitClaimBadge(c *gin.Context) {
	var req request.PermitClaimBadgeReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.PermitClaimBadge(address, req); err != nil {
		if err.Error() == "QuestUpdate" {
			OkWithMessage(GetMessage(c, err.Error()), c)
			return
		}
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(list, c)
	}
}

func SubmitClaimTweet(c *gin.Context) {
	var req request.SubmitClaimTweetReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if err := srv.SubmitClaimTweet(address, req); err != nil {
		if err.Error() == "QuestUpdate" {
			OkWithMessage(GetMessage(c, err.Error()), c)
			return
		}
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}

func UpdateBadgeURI(c *gin.Context) {
	var badgeURI request.UpdateBadgeURIRequest
	if err := c.ShouldBindJSON(&badgeURI); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.UpdateBadgeURI(address, badgeURI); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(list, c)
	}
}

func SubmitClaimShare(c *gin.Context) {
	var req request.SubmitClaimShareReq
	_ = c.ShouldBindJSON(&req)
	address := c.GetString("address")
	if res, err := srv.SubmitClaimShare(address, req); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(res, c)
	}
}
