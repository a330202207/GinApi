package backend

import (
	"GinApi/model"
	"GinApi/package/error"
	"GinApi/service"
	"GinApi/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//管理员列表
func GetAdminList(c *gin.Context) {
	admins, _ := model.GetAdminList(50, "created_at desc", map[string]interface{}{})
	c.HTML(http.StatusOK, "admin_index.html", gin.H{
		"admins": admins,
	})
}

//添加管理员页
func AdminCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_create.html", gin.H{
		"action": "add",
	})
}

//添加管理员
func AdminAdd(c *gin.Context) {
	var admin service.AdminInfo
	admin.CreateIp = c.ClientIP()

	if err := c.ShouldBind(&admin); err == nil {
		resCode := admin.AdminAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除管理员
func AdminDel(c *gin.Context) {
	var admin service.AdminInfo
	id, err := strconv.Atoi(c.PostForm("id"))
	admin.ID = id
	if id != 0 || err != nil {
		resCode := admin.AdminDel()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//编辑管理员页
func AdminEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_edit.html", gin.H{
		"action": "edit",
		"title":  "编辑",
	})
}

//保存管理员
func AdminSave(c *gin.Context) {

}
