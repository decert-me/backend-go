package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

func GetProgressList(c *gin.Context) {
	var req request.GetProgressListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	userID := c.GetUint("userID")
	if data, err := srv.GetProgressList(userID, req); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}
}

func GetProgress(c *gin.Context) {
	var req request.GetProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	userID := c.GetUint("userID")
	if userID == 0 {
		FailWithDetailed(gin.H{"reload": true}, GetMessage(c, "UnauthorizedAccess"), c)
		return
	}
	if data, err := srv.GetProgress(userID, req); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithData(data, c)
	}
}

func UpdateProgress(c *gin.Context) {
	var req request.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	userID := c.GetUint("userID")
	if userID == 0 {
		FailWithDetailed(gin.H{"reload": true}, GetMessage(c, "UnauthorizedAccess"), c)
		return
	}
	if err := srv.UpdateProgress(userID, req); err != nil {
		FailWithMessage(GetMessage(c, "UpdateFailed"), c)
	} else {
		Ok(c)
	}
}

// GetTutorialList 获取教程列表
func GetTutorialList(c *gin.Context) {
	var pageInfo request.GetTutorialListStatusRequest
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}

	if list, total, err := srv.GetTutorialList(pageInfo); err != nil {
		FailWithMessage("获取失败", c)
	} else {
		OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetLabelList 获取标签列表
func GetLabelList(c *gin.Context) {
	var label request.GetLabelRequest
	err := c.ShouldBindJSON(&label)
	if err != nil {
		FailWithMessage("参数错误", c)
		return
	}
	var data interface{}
	if label.Type == "language" {
		data, err = srv.LabelLangList()
	} else if label.Type == "category" {
		data, err = srv.LabelCategoryList(label.Class)
	} else if label.Type == "theme" {
		data, err = srv.LabelThemeList()
	} else {
		FailWithMessage("参数错误", c)
		return
	}
	if err != nil {
		FailWithMessage("获取失败", c)
	} else {
		OkWithDetailed(data, "获取成功", c)
	}
}
