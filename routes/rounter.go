package routes

import (
	api "GinApi/controller/api/v1"
	"GinApi/controller/backend"
	"GinApi/controller/frontend"
	"GinApi/middleware/logger"
	"GinApi/pkg/setting"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//初始化路由
func InitRouter(e *gin.Engine) {

	// 使用 Logger 中间件
	e.Use(gin.Logger())

	// 使用 Recovery 中间件
	e.Use(gin.Recovery())

	// 使用日志中间件
	e.Use(logger.LoggerToFile())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 模板函数
	e.SetFuncMap(template.FuncMap{
		"unescaped":   unescaped,
		"strtime":     StrTime,
		"plus1":       selfPlus,
		"numplusplus": numPlusPlus,
		"strip":       Long2IPString,
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该路由",
		})
		return
	})

	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "找不到该方法",
		})
		return
	})

	//加载模板
	loadTemplate(e)

	//注册API路由
	RegisterApiRouter(e)

	//注册前台路由
	RegisterFrontendRouter(e)

	//注册后台路由
	RegisterBackendRouter(e)
}

func loadTemplate(e *gin.Engine) {
	//加载views文件夹下所有的文件
	e.LoadHTMLGlob("views/*/**/***")

	// 推荐使用绝对路径 相当于简历了软连接--快捷方式
	e.StaticFS("/static", http.Dir("./static"))
	e.StaticFS("/upload", http.Dir("./upload"))
}

//注册API模块路由
func RegisterApiRouter(e *gin.Engine) {
	apiRouter := e.Group("/api/v1")
	{
		apiRouter.GET("/test/index", api.Test)
	}
}

func RegisterFrontendRouter(e *gin.Engine) {
	web := e.Group("")
	{
		// 首页
		web.GET("/", frontend.Index)
	}
}

func RegisterBackendRouter(e *gin.Engine) {
	admin := e.Group("/admin")
	{
		admin.GET("/", backend.Index)
		admin.GET("/login.html", backend.AdminLoginIndex)
		admin.GET("/login", backend.AdminLogin)
	}
}
