package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rechargeApi/pkg/setting"
	"rechargeApi/routes"
)

func Init() *gin.Engine {

	engine := gin.New()

	//初始化路由
	routes.InitRouter(engine)

	return engine
}

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
