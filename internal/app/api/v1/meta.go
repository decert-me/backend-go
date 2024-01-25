package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleMetaRequest(c *gin.Context) {
	tokenID := c.Param("id")
	q, _ := c.GetQuery("q")

	c.Data(http.StatusOK, "text/html; charset=utf-8", srv.HandleMetaRequest(tokenID, q))
}

func HandleCollectionMetaRequest(c *gin.Context) {
	id := c.Param("id")
	c.Data(http.StatusOK, "text/html; charset=utf-8", srv.HandleCollectionMetaRequest(id))
}
