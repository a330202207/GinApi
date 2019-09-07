package routes

import (
	"GinApi/controller/backend"
	"github.com/gin-gonic/gin"
)

//后台
func RegisterBackendRouter(e *gin.Engine) {
	admin := e.Group("/admin")
	{
		//登录页
		admin.GET("/login.html", backend.AdminLoginIndex)

		//登录
		admin.POST("/login", backend.AdminLogin)

		//首页
		admin.GET("/index.html", backend.Index)

	}
}
