package v1

import (
	"backend-go/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

func UploadJson(c *gin.Context) {
	json := c.PostForm("body")
	if json == "" {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	// 文件大小限制
	if len(json) > 1024*1024*5 {
		FailWithMessage(GetMessage(c, "FileSizeExceedsLimit"), c)
		return
	}
	err, hash := srv.IPFSUploadJSON(json) // 文件上传后拿到文件路径
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
	err, hash := srv.IPFSUploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		FailWithMessage(GetMessage(c, "UploadFailed"), c)
		return
	}
	OkWithDetailed(response.UploadResponse{Hash: hash}, GetMessage(c, "UploadSuccess"), c)
}
