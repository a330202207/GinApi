package service

import (
	"GinApi/model"
	"GinApi/package/error"
	"encoding/json"
	"fmt"
)

type Menu struct {
	ID         int    `form:"id" json:"id"`
	ParentId   int    `form:"parent_id" json:"parent_id"`
	Name       string `form:"name" json:"name" binding:"required"`
	MenuRouter string `form:"menu_router" json:"menu_router"`
	OrderBy    int    `form:"order_by" json:"order_by"`
}

//添加菜单
func (menuInfo *Menu) MenuAdd() int {
	name := map[string]interface{}{"name": menuInfo.Name}
	isExist := model.ExistMenu(name)

	if isExist == true {
		return error.ERROR_EXIST_ROLE
	}
	menu := model.Menu{
		ParentId:   menuInfo.ParentId,
		Name:       menuInfo.Name,
		MenuRouter: menuInfo.MenuRouter,
		OrderBy:    menuInfo.OrderBy,
	}
	if err := model.AddMenu(&menu); err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//删除菜单
func (menuInfo *Menu) MenuDel() int {
	id := map[string]interface{}{"id": menuInfo.ID}
	isExist := model.ExistMenu(id)
	if isExist == false {
		return error.ERROR_NOT_EXIST_RESOURCE
	}

	err := model.DelMenu(map[string]interface{}{"id": menuInfo.ID})
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}
	return error.SUCCESS
}

//编辑菜单
func (menuInfo *Menu) MenuEdit() (model.Menu, int) {
	menu, err := model.GetMenu(map[string]interface{}{"id": menuInfo.ID})
	if err != nil {
		return menu, error.ERROR
	}
	return menu, error.SUCCESS
}

//保存菜单
func (menuInfo *Menu) MenuSave() int {
	id := menuInfo.ID
	menu := model.Menu{
		ParentId:   menuInfo.ParentId,
		Name:       menuInfo.Name,
		MenuRouter: menuInfo.MenuRouter,
		OrderBy:    menuInfo.OrderBy,
	}
	if err := model.SaveMenu(id, menu); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}
	return error.SUCCESS
}

//获取菜单树
func GetTreeMenus() string {
	var menu model.Menu
	list := menu.GetTreeMenus(0)

	body, err := json.Marshal(list)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
