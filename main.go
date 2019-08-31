package main

import (
	"rechargeApi/app"
)

func main() {

	//初始化配置文件
	engine := app.Init()

	//运行服务
	app.Run(engine)

}
