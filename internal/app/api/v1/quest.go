package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

func GetQuestList(c *gin.Context) {
	var searchInfo request.GetQuestListRequest
	_ = c.ShouldBindQuery(&searchInfo)
	searchInfo.Address = c.GetString("address")
	if list, total, err := srv.GetQuestList(searchInfo); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchInfo.Page,
			PageSize: searchInfo.PageSize,
		}, GetMessage(c, "FetchSuccess"), c)
	}
}

func GetQuest(c *gin.Context) {
	if list, err := srv.GetQuest(c.Param("id"), c.GetString("address")); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(list, c)
	}
}

func AddQuest(c *gin.Context) {
	var add request.AddQuestRequest
	if err := c.ShouldBindJSON(&add); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.AddQuest(address, add); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(list, c)
	}
}
func UpdateQuest(c *gin.Context) {
	var modify request.UpdateQuestRequest
	if err := c.ShouldBindJSON(&modify); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.UpdateQuest(address, modify); err != nil {
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

// GetQuestFlashRank 获取闪电榜
func GetQuestFlashRank(c *gin.Context) {
	address := c.GetString("address")
	if data, err := srv.GetQuestFlashRank(address, c.Param("id")); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}
}

// GetQuestHighRank 获取高分榜
func GetQuestHighRank(c *gin.Context) {
	address := c.GetString("address")
	if data, err := srv.GetQuestHighRank(address, c.Param("id")); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}
}

// GetQuestHolderRank 获取持有榜
func GetQuestHolderRank(c *gin.Context) {
	var searchInfo request.GetQuestHolderRankRequest
	_ = c.ShouldBindQuery(&searchInfo)
	address := c.GetString("address")
	if data, total, err := srv.GetQuestHolderRank(address, c.Param("id"), searchInfo.Page, searchInfo.PageSize); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithDetailed(response.PageResult{
			List:     data,
			Total:    total,
			Page:     searchInfo.Page,
			PageSize: searchInfo.PageSize,
		}, GetMessage(c, "FetchSuccess"), c)
	}
}
