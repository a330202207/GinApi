package casbin

import (
	"GinApi/model"
	"GinApi/util/convert"
	"fmt"
	"github.com/casbin/casbin"
)

var Enforcer *casbin.Enforcer

//初始化 角色-URL
func InitCasbin() (err error) {
	var enforcer *casbin.Enforcer
	//文档地址
	//https://casbin.org/docs/zh-CN/function
	casbinModel := `[request_definition]
	r = sub, obj, act
	
	[policy_definition]
	p = sub, obj, act
	
	[role_definition]
	g = _, _
	
	[policy_effect]
	e = some(where (p.eft == allow))
	
	[matchers]
	m = g(r.sub, p.sub) && regexMatch(r.act, p.act)`

	//m = r.sub == p.sub && regexMatch(r.act, p.act)
	enforcer, err = casbin.NewEnforcerSafe(
		casbin.NewModel(casbinModel),
	)
	if err != nil {
		return
	}
	list, err := model.GetRoleMenus(map[string]interface{}{})
	if err != nil {
		return
	}

	for _, k := range list {
		setRolePermission(enforcer, k.RoleID, k.MenuID)
	}

	if len(list) == 0 {
		Enforcer = enforcer
		return
	}

	Enforcer = enforcer

	return
}

//设置角色权限
func setRolePermission(e *casbin.Enforcer, roleId, menuId int) {
	menu, err := model.GetMenu(map[string]interface{}{"id": menuId})
	if err != nil {
		return
	}
	e.AddPermissionForUser(convert.ToString(roleId), menu.MenuRouter, "GET|POST")
}

//检查用户是否有权限
func CheckPermission(adminId, url, methodtype string) (bool, error) {
	fmt.Println("adminId:", adminId)
	fmt.Println("url:", url)
	fmt.Println("methodtype:", methodtype)
	return Enforcer.EnforceSafe(adminId, url, methodtype)
}

//用户角色处理
func AddRoleForUser(adminId int) (err error) {
	if Enforcer == nil {
		return
	}
	Enforcer.DeleteRolesForUser(convert.ToString(adminId))

	list, err := model.GetAdminRoles(map[string]interface{}{"admin_id": adminId})
	if err != nil {
		return
	}
	for _, v := range list {
		Enforcer.AddRoleForUser(convert.ToString(adminId), convert.ToString(v.RoleID))
	}
	return
}
