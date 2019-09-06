package routes

import (
	api "GinApi/controller/api/v1"
	"GinApi/controller/backend"
	"GinApi/controller/frontend"
	"GinApi/middleware/cors"
	"GinApi/middleware/error"
	"GinApi/middleware/logger"
	"GinApi/middleware/session"
	"GinApi/pkg/setting"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//初始化路由
func InitRouter(e *gin.Engine) {
	//session
	e.Use(session.Session())

	// 使用 Logger 中间件
	e.Use(gin.Logger())

	// 使用 Recovery 中间件
	e.Use(gin.Recovery())

	// 使用日志中间件
	e.Use(logger.LoggerToFile())

	//跨域
	e.Use(cors.Cors())

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
			"code": error.NOROUTE,
			"msg":  error.GetMsg(error.NOROUTE),
		})
		return
	})

	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": error.NOROUTE,
			"msg":  error.GetMsg(error.NOROUTE),
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

//前台
func RegisterFrontendRouter(e *gin.Engine) {
	web := e.Group("")
	{
		// 首页
		web.GET("/", frontend.Index)
	}
}

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
