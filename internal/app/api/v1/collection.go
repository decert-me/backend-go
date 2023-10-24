package v1

import (
	"backend-go/internal/app/model/request"
	"backend-go/internal/app/model/response"
	"github.com/gin-gonic/gin"
)

// GetCollectionQuest 获取合辑内挑战
func GetCollectionQuest(c *gin.Context) {
	var r request.GetCollectionQuestRequest
	_ = c.ShouldBindQuery(&r)
	r.Address = c.GetString("address")
	if list, collection, err := srv.GetCollectionQuest(r); err != nil {
		FailWithMessage(GetMessage(c, "FetchFailed"), c)
	} else {
		OkWithDetailed(response.GetCollectionQuestPageResult{
			List:       list,
			Collection: collection,
		}, GetMessage(c, "FetchSuccess"), c)
	}
}

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

// CollectionClaim 领取合辑奖励
func CollectionClaim(c *gin.Context) {
	var r request.CollectionClaimRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		FailWithMessage("ParameterError", c)
		return
	}
	if err := srv.CollectionClaim(r, c.GetString("address")); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}

// CheckQuestInCollection 查询挑战是否在合辑内
func CheckQuestInCollection(c *gin.Context) {
	var r request.CheckQuestInCollectionRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		FailWithMessage("ParameterError", c)
		return
	}
	if res, err := srv.CheckQuestInCollection(r); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(res, c)
	}
}
