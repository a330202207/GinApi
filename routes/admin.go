package routes

import (
	"GinApi/controller/backend"
	"github.com/gin-gonic/gin"
)

//后台
func RegisterBackendRouter(e *gin.Engine) {
	admin := e.Group("/admin")
	{
		//登录页
		admin.GET("/backend_login.html", backend.AdminLoginIndex)

		//登录
		admin.POST("/login", backend.AdminLogin)

		//登出
		admin.GET("/logout", backend.AdminLogOut)

		//首页
		admin.GET("/backend_index.html", backend.Index)

		//管理员列表
		admin.GET("/admin/admin_index.html", backend.GetAdminList)

		//添加管理员页
		admin.GET("/admin/admin_create.html", backend.AdminCreate)

		//添加管理员
		admin.POST("/admin/add", backend.AdminAdd)

		//删除管理员
		admin.POST("/admin/del", backend.AdminDel)

		//编辑管理员页
		admin.GET("/admin/admin_edit.html", backend.AdminEdit)

		//保存管理员信息
		admin.POST("/admin/save", backend.AdminSave)

		//修改管理员密码页
		admin.GET("/admin/edit_password.html", backend.AdminEdit)

		//角色列表
		admin.GET("/role/role_index.html", backend.GetRoleList)

		//添加角色页
		admin.GET("/role/role_create.html", backend.RoleCreate)

		//添加角色
		admin.POST("/role/add", backend.RoleAdd)

		//删除角色
		admin.POST("/role/del", backend.RoleDel)

		//编辑角色页
		admin.GET("/role/role_edit.html", backend.RoleEdit)

		//获取当前角色菜单
		admin.GET("/role/myMenus", backend.MyMenus)

		//保存角色
		admin.POST("/role/save", backend.RoleSave)

		//菜单列表页
		admin.GET("/menu/menu_index.html", backend.GetMenuList)

		//获取菜单树结构
		admin.GET("/menu/menus", backend.GetTreeMenus)

		//添加菜单页
		admin.GET("/menu/menu_create.html", backend.MenuCreate)

		//添加菜单
		admin.POST("/menu/add", backend.MenuAdd)

		//添加下级菜单页
		admin.GET("/menu/menu_add.html", backend.AddMenu)

		//删除菜单
		admin.POST("/menu/del", backend.MenuDel)

		//编辑菜单页
		admin.GET("/menu/menu_edit.html", backend.MenuEdit)

		//保存菜单
		admin.POST("/menu/save", backend.MenuSave)

	}
}
