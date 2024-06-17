package v2

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"

	"github.com/gin-gonic/gin"
)

// GetUserOpenQuestList 获取用户开放题列表
func GetUserOpenQuestList(c *gin.Context) {
	var r request.GetUserOpenQuestListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		FailWithMessage("参数错误", c)
		return
	}
	address := c.GetString("address")
	type Detail struct {
		List          interface{} `json:"list"`
		Total         int64       `json:"total"`
		Page          int         `json:"page"`
		PageSize      int         `json:"pageSize"`
		TotalToReview int64       `json:"total_to_review"`
	}
	if list, total, totalToReview, err := srv.GetUserOpenQuestListV2(address, r); err != nil {
		FailWithMessage("获取失败："+err.Error(), c)
	} else {
		OkWithDetailed(Detail{
			List:          list,
			Total:         total,
			Page:          r.Page,
			PageSize:      r.PageSize,
			TotalToReview: totalToReview,
		}, "获取成功", c)
	}
}

// GetUserOpenQuestDetailList 获取开放题待评分列表
func GetUserOpenQuestDetailList(c *gin.Context) {
	var r request.GetUserOpenQuestDetailListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		FailWithMessage("参数错误"+err.Error(), c)
		return
	}
	address := c.GetString("address")
	LoginAddress := c.GetString("address")
	if list, total, err := srv.GetUserOpenQuestDetailListV2(address, LoginAddress, r); err != nil {
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
