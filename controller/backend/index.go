package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	//sessions := models.GetSessions(c)

	title := "Hello Admin!"
	c.HTML(http.StatusOK, "admin_index.html", gin.H{
		"title": title,
	})
}
