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
	if list, err := srv.GetQuest(c.Param("id")); err != nil {
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
