package backend

import (
	"GinApi/config"
	"GinApi/model"
	"GinApi/package/error"
	"GinApi/service"
	"GinApi/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

//角色列表页
func GetRoleList(c *gin.Context) {

	data := map[string]interface{}{"status <": 2}
	name := c.Query("keyword")
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize := config.ServerSetting.PageSize
	Offset := (page - 1) * pageSize

	if name != "" {
		data["name like"] = "%" + name
	}

	query, args, _ := util.WhereBuild(data)
	fmt.Println(query, args)
	roles, count, _ := model.GetRoleList(pageSize, Offset, "created_at desc", query, args...)

	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	pagination := util.NewPagination(c.Request, count, pageSize)

	c.HTML(http.StatusOK, "role_index.html", gin.H{
		"roles":     roles,
		"name":      name,
		"totalPage": totalPage,
		"count":     count,
		"pages":     template.HTML(pagination.Pages()),
	})
}

//添加角色页
func RoleCreate(c *gin.Context) {
	c.HTML(http.StatusOK, "role_create.html", gin.H{
		"action": "add",
	})
}

//添加角色
func RoleAdd(c *gin.Context) {
	var role service.RoleInfo
	fmt.Println(role)
	if err := c.ShouldBind(&role); err == nil {
		resCode := role.RoleAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}

}

//删除角色
func RoleDel(c *gin.Context) {
	var role service.RoleInfo
	id, err := strconv.Atoi(c.PostForm("id"))
	role.ID = id
	if id != 0 || err != nil {
		resCode := role.RoleDel()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//编辑角色页
func RoleEdit(c *gin.Context) {
	var role service.RoleInfo
	id, err := strconv.Atoi(c.Query("id"))
	role.ID = id
	if id != 0 || err != nil {
		if info, errCode := role.RoleEdit(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			//登陆成功
			c.HTML(http.StatusOK, "role_edit.html", gin.H{
				"action": "edit",
				"title":  "编辑",
				"info":   info,
			})
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//保存角色
func RoleSave(c *gin.Context) {
	var role service.RoleInfo
	if err := c.ShouldBind(&role); err == nil {
		resCode := role.RoleSave()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}
