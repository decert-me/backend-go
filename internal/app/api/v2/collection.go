package v2

import (
	"backend-go/internal/app/model/request"
	"github.com/gin-gonic/gin"
)

// CollectionClaimV2 领取合辑奖励
func CollectionClaimV2(c *gin.Context) {
	var r request.CollectionClaimRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		FailWithMessage(GetMessage(c, "ParameterError"), c)
		return
	}
	if err := srv.CollectionClaimV2(r, c.GetString("address")); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		Ok(c)
	}
}
