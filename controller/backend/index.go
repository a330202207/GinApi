package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//首页
func Index(c *gin.Context) {
	//sessions := models.GetSessions(c)

	title := "Hello Admin!"
	c.HTML(http.StatusOK, "backend_index.html", gin.H{
		"title": title,
	})
}
