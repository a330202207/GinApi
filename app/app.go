package app

import (
	"GinApi/models"
	"GinApi/pkg/setting"
	"GinApi/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//初始化
func Init() *gin.Engine {

	engine := gin.New()

	//初始化路由
	routes.InitRouter(engine)

	//加载服务
	LoadServer()

	return engine
}

//加载服务
func LoadServer() {
	//加载数据库
	models.Database()
}

//运行
func Run(e *gin.Engine) {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        e,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
