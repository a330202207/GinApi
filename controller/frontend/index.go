package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	title := "Hello World"
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": title,
	})
}
