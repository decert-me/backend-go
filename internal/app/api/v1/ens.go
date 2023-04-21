package v1

import "github.com/gin-gonic/gin"

func GetEnsRecords(c *gin.Context) {
	if list, err := srv.GetEnsRecords(c, c.Param("q")); err != nil {
		FailWithMessage(GetMessage(c, err.Error()), c)
	} else {
		OkWithData(list, c)
	}
}
