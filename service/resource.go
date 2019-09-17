package service

import (
	"GinApi/model"
	"GinApi/package/error"
	"encoding/json"
	"fmt"
)

type Resource struct {
	ID             int    `form:"id" json:"id"`
	PID            int    `form:"p_id" json:"p_id"`
	Name           string `form:"name" json:"name" binding:"required"`
	ResourceRouter string `form:"resource_router" json:"resource_router"`
}

//添加资源
func (resourceInfo *Resource) ResourceAdd() int {
	name := map[string]interface{}{"name": resourceInfo.Name}
	isExist := model.ExistResource(name)

	if isExist == true {
		return error.ERROR_EXIST_ROLE
	}
	resource := model.Resource{
		PID:            resourceInfo.PID,
		Name:           resourceInfo.Name,
		ResourceRouter: resourceInfo.ResourceRouter,
	}
	if err := model.AddResource(&resource); err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//删除资源
func (resourceInfo *Resource) ResourceDel() int {
	id := map[string]interface{}{"id": resourceInfo.ID}
	isExist := model.ExistResource(id)
	if isExist == false {
		return error.ERROR_NOT_EXIST_RESOURCE
	}

	err := model.DelResource(map[string]interface{}{"id": resourceInfo.ID})
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}
	return error.SUCCESS
}

//编辑资源
func (resourceInfo *Resource) ResourceEdit() (model.Resource, int) {
	resource, err := model.GetResource(map[string]interface{}{"id": resourceInfo.ID})
	if err != nil {
		return resource, error.ERROR
	}
	return resource, error.SUCCESS
}

//保存资源
func (resourceInfo *Resource) ResourceSave() int {
	id := resourceInfo.ID
	resource := model.Resource{
		PID:            resourceInfo.PID,
		Name:           resourceInfo.Name,
		ResourceRouter: resourceInfo.ResourceRouter,
	}
	if err := model.SaveResource(id, resource); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}
	return error.SUCCESS
}

//获取资源树
func GetTreeResources() string {
	var resource model.Resource
	list := resource.GetTreeResources(0)

	body, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
