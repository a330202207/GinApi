package service

import (
	"GinApi/model"
	"GinApi/package/error"
)

type RoleInfo struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
}

//添加角色
func (roleInfo *RoleInfo) RoleAdd() int {
	name := map[string]interface{}{"name": roleInfo.Name, "status": 1}
	isExist := model.ExistRole(name)

	if isExist == true {
		return error.ERROR_EXIST_ROLE
	}
	role := model.Role{
		Name:   roleInfo.Name,
		Status: 1,
	}

	if err := model.AddRole(&role); err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}
	return error.SUCCESS
}

//删除角色
func (role *RoleInfo) RoleDel() int {
	userName := map[string]interface{}{"id": role.ID, "status": 1}
	isExist := model.ExistRole(userName)
	if isExist == false {
		return error.ERROR_NOT_EXIST_USER
	}

	err := model.DelRole(map[string]interface{}{"id": role.ID})
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}
	return error.SUCCESS
}

//编辑角色
func (role *RoleInfo) RoleEdit() (model.Role, int) {
	roleInfo, err := model.GetRole(map[string]interface{}{"id": role.ID})
	if err != nil {
		return roleInfo, error.ERROR
	}
	return roleInfo, error.SUCCESS
}

//保存角色
func (roleInfo *RoleInfo) RoleSave() int {
	id := roleInfo.ID
	role := model.Role{
		Name: roleInfo.Name,
	}
	if err := model.SaveRole(id, role); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}
	return error.SUCCESS
}
