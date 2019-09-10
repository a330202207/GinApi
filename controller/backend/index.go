package backend

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//首页
func Index(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	fmt.Println(userId)
	c.HTML(http.StatusOK, "backend_index.html", gin.H{
		"user_id": userId,
	})
}
