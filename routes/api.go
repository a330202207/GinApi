package routes

import (
	api "GinApi/controller/api/v1"
	"github.com/gin-gonic/gin"
)

//注册API模块路由
func RegisterApiRouter(e *gin.Engine) {
	apiRouter := e.Group("/api/v1")
	{
		apiRouter.GET("/test/index", api.Test)
	}
}
