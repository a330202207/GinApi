package main

import (
	"GinApi/app"
)

func main() {

	//初始化
	engine := app.Init()

	//运行服务
	app.Run(engine)

}
