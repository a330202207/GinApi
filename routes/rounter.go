package routes

import (
	"GinApi/config"
	"GinApi/middleware/casbin"
	"GinApi/middleware/cors"
	"GinApi/middleware/session"
	"GinApi/package/error"
	"GinApi/util"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

//初始化路由
func InitRouter(e *gin.Engine) {
	//session
	e.Use(session.Session())

	//使用 Logger 中间件
	e.Use(gin.Logger())

	//使用 Recovery 中间件
	e.Use(gin.Recovery())

	//使用日志中间件
	//e.Use(logger.LoggerToFile())

	//权限中间件
	e.Use(casbin.CheckLoginHandle(
		casbin.AllowPathPrefixSkipper(casbin.NotCheckPermissionUrl()...),
	))

	//跨域
	e.Use(cors.Cors())

	//设置环境
	gin.SetMode(config.ServerSetting.RunMode)

	//模板函数
	e.SetFuncMap(template.FuncMap{
		"IntToTime": util.IntToTime,
	})

	//404
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

//加载模板
func loadTemplate(e *gin.Engine) {
	//加载views文件夹下所有的文件
	e.LoadHTMLGlob("views/*/**/***")

	e.StaticFS("/static", http.Dir("./static"))
	e.StaticFS("/upload", http.Dir("./upload"))
	e.StaticFile("favicon.ico", "/static/favicon.ico")

}
