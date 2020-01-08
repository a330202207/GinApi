package service

import (
	"GinApi/model"
	"GinApi/package/error"
)

type RoleInfo struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
}

type RoleMenu struct {
	RoleInfo
	MenuIDs []string `form:"menu_ids" json:"menu_ids" binding:"required"`
}

//添加角色
func (roleInfo *RoleMenu) RoleAdd() int {
	name := map[string]interface{}{"name": roleInfo.Name, "status": 1}
	isExist := model.ExistRole(name)

	if isExist == true {
		return error.ERROR_EXIST_ROLE
	}
	role := model.Role{
		Name:   roleInfo.Name,
		Status: 1,
	}

	roleID, err := model.AddRole(&role)
	if err != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}

	if model.AddRoleMenu(roleID, roleInfo.MenuIDs) != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}

	return error.SUCCESS
}

//删除角色
func (role *RoleInfo) RoleDel() int {
	name := map[string]interface{}{"id": role.ID, "status": 1}
	isExist := model.ExistRole(name)
	if isExist == false {
		return error.ERROR_NOT_EXIST_USER
	}

	err := model.DelRole(map[string]interface{}{"id": role.ID})
	if err != nil {
		return error.ERROR_SQL_DELETE_FAIL
	}

	if model.DelRoleMenu(role.ID) != nil {
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
func (roleInfo *RoleMenu) RoleSave() int {
	id := roleInfo.ID
	role := model.Role{
		Name: roleInfo.Name,
	}
	if err := model.SaveRole(id, role); err != nil {
		return error.ERROR_SQL_UPDATE_FAIL
	}

	if model.DelRoleMenu(roleInfo.RoleInfo.ID) != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}

	if model.AddRoleMenu(roleInfo.RoleInfo.ID, roleInfo.MenuIDs) != nil {
		return error.ERROR_SQL_INSERT_FAIL
	}

	//添加权限

	return error.SUCCESS
}
