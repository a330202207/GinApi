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

//资源列表页
func GetResourceList(c *gin.Context) {
	data := map[string]interface{}{}
	name := c.Query("keyword")
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize := config.ServerSetting.PageSize

	Offset := (page - 1) * pageSize

	if name != "" {
		data["name like"] = name + "%"
	}

	query, args, _ := util.WhereBuild(data)
	Resources, count, _ := model.GetResourceList(pageSize, Offset, "id asc", query, args...)

	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	pagination := util.NewPagination(c.Request, count, pageSize)

	c.HTML(http.StatusOK, "resource_index.html", gin.H{
		"resources": Resources,
		"totalPage": totalPage,
		"count":     count,
		"pages":     template.HTML(pagination.Pages()),
	})
}

//添加资源页
func ResourceCreate(c *gin.Context) {
	topResources, _ := model.GetResources(map[string]interface{}{"p_id": 0})
	c.HTML(http.StatusOK, "resource_create.html", gin.H{
		"resources": topResources,
	})
}

//添加资源
func ResourceAdd(c *gin.Context) {
	var resource service.Resource
	if err := c.ShouldBind(&resource); err == nil {
		resCode := resource.ResourceAdd()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//删除资源
func ResourceDel(c *gin.Context) {
	var resource service.Resource
	id, err := strconv.Atoi(c.PostForm("id"))
	resource.ID = id
	if id != 0 || err != nil {
		resCode := resource.ResourceDel()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//编辑资源页
func ResourceEdit(c *gin.Context) {
	var resource service.Resource
	id, err := strconv.Atoi(c.Query("id"))
	resource.ID = id
	if id != 0 || err != nil {
		if info, errCode := resource.ResourceEdit(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			topResources, _ := model.GetResources(map[string]interface{}{"p_id": 0})
			c.HTML(http.StatusOK, "resource_edit.html", gin.H{
				"action":    "edit",
				"title":     "编辑",
				"info":      info,
				"resources": topResources,
			})
		}
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//保存资源
func ResourceSave(c *gin.Context) {
	var resource service.Resource
	if err := c.ShouldBind(&resource); err == nil {
		resCode := resource.ResourceSave()
		util.HtmlResponse(c, resCode)
	} else {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
	}
}

//获取资源树
func GetTreeResources(c *gin.Context) {

	list := service.GetTreeResources()

	util.JsonSuccessResponse(c, error.SUCCESS, list)
}
