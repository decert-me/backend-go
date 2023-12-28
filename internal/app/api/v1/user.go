package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/utils"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

func UpdateAvatar(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		FailWithMessage("ReceiveFileFailed", c)
		return
	}
	// 文件大小限制
	if header.Size > 1024*1024*4 {
		FailWithMessage(GetMessage(c, "FileSizeExceedsLimit"), c)
		return
	}
	// 读取文件后缀
	ext := strings.ToLower(path.Ext(header.Filename))
	// 限制文件后缀
	if (ext == ".jpg" || ext == ".png" || ext == ".jpeg" || ext == ".gif") == false {
		FailWithMessage(GetMessage(c, "FileFormatIncorrect"), c)
		return
	}
	if c.GetString("address") == "" {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	if p, err := srv.UpdateAvatar(c.GetString("address"), header); err != nil {
		FailWithMessage(GetMessage(c, "UpdateFailed"), c)
	} else {
		OkWithData(map[string]string{"url": p}, c)
	}
}

func GetDiscordInfo(c *gin.Context) {
	if list, err := srv.GetDiscordInfo(c.GetString("address")); err != nil {
		FailWithMessage(GetMessage(c, "NoBindingDetected"), c)
	} else {
		OkWithData(list, c)
	}
}

func GetTwitterInfo(c *gin.Context) {
	if list, err := srv.GetTwitterInfo(c.GetString("address")); err != nil {
		FailWithMessage(GetMessage(c, "NoBindingDetected"), c)
	} else {
		OkWithData(list, c)
	}
}

// GetUserInfo Get User Info
func GetUserInfo(c *gin.Context) {
	address := c.Param("address")
	if address == "" {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	if list, err := srv.GetUserInfo(address); err != nil {
		OkWithData(list, c)
	} else {
		OkWithData(list, c)
	}
}

// UpdateUserInfo update User Info
func UpdateUserInfo(c *gin.Context) {
	var request request.UpdateUserInfo
	_ = c.ShouldBindJSON(&request)
	address := c.Param("address")
	if address == "" || c.GetString("address") != address {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	if err := srv.UpdateUserInfo(address, request); err != nil {
		FailWithMessage(GetMessage(c, "UpdateFailed"), c)
	} else {
		OkWithMessage(GetMessage(c, "UpdateSuccess"), c)
	}
}

func GetUserQuestList(c *gin.Context) {
	var searchInfo request.GetUserQuestListRequest
	_ = c.ShouldBindQuery(&searchInfo)
	searchInfo.Creator = c.Param("addressCreator")
	searchInfo.Address = c.GetString("address")
	searchInfo.Language = c.GetString("lang")
	if list, total, err := srv.GetUserQuestListWithClaimed(searchInfo); err != nil {
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

// GetUserChallengeList Get User Challenge list
func GetUserChallengeList(c *gin.Context) {
	var searchInfo request.GetChallengeListRequest
	_ = c.ShouldBindQuery(&searchInfo)
	address := c.Param("address")
	if address == "" {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	searchInfo.ReqAddress = c.GetString("address")
	searchInfo.Address = address
	searchInfo.Language = c.GetString("lang")
	if list, total, err := srv.GetUserChallengeList(searchInfo); err != nil {
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

// GetLoginMessage
// @Tags SignApi
// @Summary 获取登录签名消息
// @accept application/json
// @Produce application/json
// @Router /sign/getLoginMessage [post]
func GetLoginMessage(c *gin.Context) {
	var request request.GetLoginMessageRequest
	_ = c.ShouldBindQuery(&request)
	if loginMessage, err := srv.GetLoginMessage(request.Address); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithDetailed(map[string]string{"loginMessage": loginMessage}, GetMessage(c, "FetchSuccess"), c)
	}
}

// AuthLoginSign
// @Tags SignApi
// @Summary 校验登录签名
// @accept application/json
// @Produce application/json
// @Router /sign/authLoginSign [post]
func AuthLoginSign(c *gin.Context) {
	var request request.AuthLoginSignRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	var token string
	var err error
	if utils.IsValidAddress(request.Address) {
		token, err = srv.AuthLoginSignRequest(request)
	} else {
		token, err = srv.AuthLoginSignRequestSolana(request)
	}
	if err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithDetailed(map[string]string{"token": token}, GetMessage(c, "FetchSuccess"), c)
	}
}

// HasCreateOpenQuestPerm 获取用户是否有创建开放题权限
func HasCreateOpenQuestPerm(c *gin.Context) {
	if perm, beta, err := srv.HasOpenQuestPerm(c.GetString("address")); err != nil {
		Fail(c)
	} else {
		OkWithData(map[string]interface{}{"perm": perm, "beta": beta}, c)
	}
}

// HasBindSocialAccount 获取用户是否绑定社交账号
func HasBindSocialAccount(c *gin.Context) {
	if wechat, discord, err := srv.HasBindSocialAccount(c.GetString("address")); err != nil {
		Fail(c)
	} else {
		OkWithData(map[string]interface{}{"wechat": wechat, "discord": discord}, c)
	}
}
