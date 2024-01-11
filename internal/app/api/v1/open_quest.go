package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

// ReviewOpenQuest 审核开放题目V2
func ReviewOpenQuest(c *gin.Context) {
	var r []request.ReviewOpenQuestRequest
	_ = c.ShouldBindJSON(&r)
	address := c.GetString("address")
	if err := srv.ReviewOpenQuest(address, r); err != nil {
		FailWithMessage("操作失败："+err.Error(), c)
	} else {
		OkWithMessage("操作成功", c)
	}
}

// GetUserOpenQuestList 获取用户开放题列表
func GetUserOpenQuestList(c *gin.Context) {
	var r request.GetUserOpenQuestListRequest
	_ = c.ShouldBindJSON(&r)
	address := c.GetString("address")
	if list, total, err := srv.GetUserOpenQuestList(address, r); err != nil {
		FailWithMessage("获取失败："+err.Error(), c)
	} else {
		OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     r.Page,
			PageSize: r.PageSize,
		}, "获取成功", c)
	}

}
