package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//路由列表页
func GetRouterList(c *gin.Context) {

	c.HTML(http.StatusOK, "router_index.html", gin.H{
		"title": "hello",
	})
}

//添加路由页
func RouterCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "router_create.html", gin.H{
		"title": "hello",
	})
}

//添加路由
func RouterAdd(c *gin.Context) {

}

//删除路由
func RouterDel(c *gin.Context) {

}

//编辑路由页
func RouterEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "router_edit.html", gin.H{
		"title": "hello",
	})
}

//保存路由
func RouterSave(c *gin.Context) {

}
