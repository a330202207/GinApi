package backend

import (
	"GinApi/config"
	"GinApi/model"
	"GinApi/package/error"
	"GinApi/service"
	"GinApi/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

//管理员列表
func GetAdminList(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")

	data := map[string]interface{}{"status <": 3}
	name := c.Query("keyword")
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize := config.ServerSetting.PageSize
	Offset := (page - 1) * pageSize

	if name != "" {
		data["user_name like"] = name + "%"
	}

	query, args, _ := util.WhereBuild(data)
	admins, count, _ := model.GetAdminList(pageSize, Offset, "created_at desc", query, args...)

	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	pagination := util.NewPagination(c.Request, count, pageSize)

	c.HTML(http.StatusOK, "admin_index.html", gin.H{
		"userId":    userId,
		"admins":    admins,
		"username":  name,
		"totalPage": totalPage,
		"count":     count,
		"pages":     template.HTML(pagination.Pages()),
	})
}

//添加管理员页
func AdminCreate(c *gin.Context) {
	roles, _ := model.GetAllRoles()

	c.HTML(http.StatusOK, "admin_create.html", gin.H{
		"action": "add",
		"roles":  roles,
	})
}

//添加管理员
func AdminAdd(c *gin.Context) {
	var admin service.Admin
	admin.AdminInfo.CreateIp = c.ClientIP()
	if err := c.ShouldBind(&admin); err == nil {
		resCode := admin.AdminAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除管理员
func AdminDel(c *gin.Context) {
	var adminInfo service.AdminInfo
	id, err := strconv.Atoi(c.PostForm("id"))
	adminInfo.ID = id
	if id != 0 || err != nil {
		resCode := adminInfo.AdminDel()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//编辑管理员页
func AdminEdit(c *gin.Context) {
	var account service.Account
	id, err := strconv.Atoi(c.Query("id"))
	account.AdminInfo.ID = id

	roles, _ := model.GetAllRoles()

	myRoles, _ := model.GetAdminRoles(map[string]interface{}{"admin_id": id})

	if id != 0 || err != nil {
		if info, errCode := account.AdminEdit(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			c.HTML(http.StatusOK, "admin_edit.html", gin.H{
				"action":  "edit",
				"title":   "编辑",
				"info":    info,
				"roles":   roles,
				"myRoles": myRoles,
			})
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}

}

//保存管理员
func AdminSave(c *gin.Context) {
	var admin service.Account
	if err := c.ShouldBind(&admin); err == nil {
		resCode := admin.AdminSave()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}
