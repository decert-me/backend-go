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
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchInfo.Page,
			PageSize: searchInfo.PageSize,
		}, "Success", c)
	}
}

func GetQuest(c *gin.Context) {
	if list, err := srv.GetQuest(c.Param("id")); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(list, c)
	}
}

func AddQuest(c *gin.Context) {
	var add request.AddQuestRequest
	if err := c.ShouldBindJSON(&add); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	address := c.GetString("address")
	if list, err := srv.AddQuest(address, add); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(list, c)
	}
}
