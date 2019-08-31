package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"mes":  "ok",
		"data": "",
	})
}
