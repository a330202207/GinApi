package backend

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//首页
func Index(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	c.HTML(http.StatusOK, "backend_index.html", gin.H{
		"user_id": userId,
	})
}
