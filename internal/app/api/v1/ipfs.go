package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"backend-go/internal/app/utils"
	"github.com/gin-gonic/gin"
)

func UploadJson(c *gin.Context) {
	types := c.Query("type")
	var uploadJSON interface{}
	if types == "challenge" {
		var uploadJSONChallenge request.UploadJSONChallenge
		if err := c.ShouldBindJSON(&uploadJSONChallenge); err != nil {
			FailWithMessage(GetMessage(c, "ParameterError"), c)
			return
		}
		if !utils.VerifyUploadJSONChallenge(srv.GetConfig().Quest.EncryptKey, uploadJSONChallenge) {
			FailWithMessage(GetMessage(c, "ParameterError"), c)
			return
		}
		uploadJSON = uploadJSONChallenge
	} else if types == "nft" {
		var uploadJSONNFT request.UploadJSONNFT
		if err := c.ShouldBindJSON(&uploadJSONNFT); err != nil {
			FailWithMessage(GetMessage(c, "ParameterError"), c)
			return
		}
		uploadJSON = uploadJSONNFT
	} else {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	err, hash := srv.IPFSUploadJSON(uploadJSON) // 文件上传后拿到文件路径
	if err != nil {
		FailWithMessage(GetMessage(c, "UploadFailed"), c)
		return
	}
	OkWithDetailed(response.UploadResponse{Hash: hash}, GetMessage(c, "UploadSuccess"), c)
}

func UploadFile(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		FailWithMessage(GetMessage(c, "ReceiveFileFailed"), c)
		return
	}
	// 文件大小限制
	if header.Size > 1024*1024*20 {
		FailWithMessage(GetMessage(c, "FileSizeExceedsLimit"), c)
		return
	}
	// 文件格式限制
	if !utils.VerifyFileFormat(header.Filename, []string{"jpeg", "jpg", "png", "gif", "svg"}) {
		FailWithMessage(GetMessage(c, "FileFormatIncorrect"), c)
		return
	}
	err, hash := srv.IPFSUploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		FailWithMessage(GetMessage(c, "UploadFailed"), c)
		return
	}
	OkWithDetailed(response.UploadResponse{Hash: hash}, GetMessage(c, "UploadSuccess"), c)
}
