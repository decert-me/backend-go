package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

// GetCollectionChallengeUser 获取挑战合辑用户
func GetCollectionChallengeUser(c *gin.Context) {
	var r request.GetCollectionChallengeUser
	if err := c.ShouldBindQuery(&r); err != nil {
		FailWithMessage("ParameterError", c)
		return
	}
	if data, total, err := srv.GetCollectionChallengeUser(r); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithDetailed(response.GetCollectionChallengeUserPageDataResult{
			GetCollectionChallengeUserRes: response.GetCollectionChallengeUserRes{Users: data.Users, Times: data.Times},
			Total:                         total,
			Page:                          r.Page,
			PageSize:                      r.PageSize,
		}, GetMessage(c, "FetchSuccess"), c)
	}
}
