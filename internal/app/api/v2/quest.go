package v2

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

func AddQuestV2(c *gin.Context) {
	var add request.AddQuestV2Request
	if err := c.ShouldBindJSON(&add); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.AddQuestV2(address, add); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(list, c)
	}
}

func UpdateQuestV2(c *gin.Context) {
	var modify request.UpdateQuestV2Request
	if err := c.ShouldBindJSON(&modify); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.UpdateQuestV2(address, modify); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(list, c)
	}
}

func GetQuestChallengeUser(c *gin.Context) {
	if data, err := srv.GetQuestChallengeUser(c.Param("id")); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}
}

func UpdateRecommend(c *gin.Context) {
	var modify request.UpdateRecommendRequest
	if err := c.ShouldBindJSON(&modify); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if err := srv.UpdateRecommend(address, modify); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}
