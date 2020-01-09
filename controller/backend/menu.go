package backend

import (
	"GinApi/config"
	"GinApi/model"
	"GinApi/package/error"
	"GinApi/service"
	"GinApi/util"
	"github.com/gin-gonic/gin"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

//菜单列表页
func GetMenuList(c *gin.Context) {
	data := map[string]interface{}{}
	name := c.Query("keyword")
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize := config.ServerSetting.PageSize

	Offset := (page - 1) * pageSize

	if name != "" {
		data["name like"] = name + "%"
	}

	query, args, _ := util.WhereBuild(data)
	Menus, count, _ := model.GetMenuList(pageSize, Offset, "id asc", query, args...)

	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	pagination := util.NewPagination(c.Request, count, pageSize)

	c.HTML(http.StatusOK, "menu_index.html", gin.H{
		"menus":     Menus,
		"totalPage": totalPage,
		"count":     count,
		"pages":     template.HTML(pagination.Pages()),
	})
}

//添加菜单页
func MenuCreate(c *gin.Context) {
	topMenus, _ := model.GetMenus(map[string]interface{}{"p_id": 0})
	c.HTML(http.StatusOK, "menu_create.html", gin.H{
		"menus": topMenus,
	})
}

//添加菜单
func MenuAdd(c *gin.Context) {
	var menu service.Menu
	if err := c.ShouldBind(&menu); err == nil {
		resCode := menu.MenuAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除菜单
func MenuDel(c *gin.Context) {
	var menu service.Menu
	id, err := strconv.Atoi(c.PostForm("id"))
	menu.ID = id
	if id != 0 || err != nil {
		resCode := menu.MenuDel()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//编辑菜单页
func MenuEdit(c *gin.Context) {
	var menu service.Menu
	id, err := strconv.Atoi(c.Query("id"))
	menu.ID = id
	if id != 0 || err != nil {
		if info, errCode := menu.MenuEdit(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {

			topMenu, _ := model.GetMenu(map[string]interface{}{"id": info.ParentId})
			c.HTML(http.StatusOK, "menu_edit.html", gin.H{
				"action": "edit",
				"info":   info,
				"menu":   topMenu,
			})
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

func AddMenu(c *gin.Context) {
	var menu service.Menu
	id, err := strconv.Atoi(c.Query("id"))

	menu.ID = id
	if id != 0 || err != nil {
		if info, errCode := menu.MenuEdit(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {

			c.HTML(http.StatusOK, "menu_add.html", gin.H{
				"action": "edit",
				"info":   info,
			})
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//保存菜单
func MenuSave(c *gin.Context) {
	var menu service.Menu
	if err := c.ShouldBind(&menu); err == nil {
		resCode := menu.MenuSave()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//获取菜单树
func GetTreeMenus(c *gin.Context) {
	list := service.GetTreeMenus()
	util.JsonSuccessResponse(c, error.SUCCESS, list)
}
