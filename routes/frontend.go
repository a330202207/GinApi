package routes

import (
	"GinApi/controller/frontend"
	"github.com/gin-gonic/gin"
)

//前台
func RegisterFrontendRouter(e *gin.Engine) {
	web := e.Group("")
	{
		// 首页
		web.GET("/", frontend.Index)
	}
}
