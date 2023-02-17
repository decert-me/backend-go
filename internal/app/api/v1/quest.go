package v1

import (
	"backend-go/internal/app/global"
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetQuestList(c *gin.Context) {
	var searchInfo request.GetQuestListRequest
	_ = c.ShouldBindQuery(&searchInfo)
	if list, err := service.GetQuestList(searchInfo); err != nil {
		global.LOG.Error("获取失败!", zap.Error(err))
		c.JSON(http.StatusOK,
			nil,
		)
	} else {
		c.JSON(http.StatusOK,
			list,
		)
	}
}
