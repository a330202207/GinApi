package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	title := "Hello Admin!"
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": title,
	})
}
