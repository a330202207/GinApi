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
		admin.GET("/login.html", backend.AdminLoginIndex)

		//登录
		admin.POST("/login", backend.AdminLogin)

		//首页
		admin.GET("/backend_index.html", backend.Index)

		//管理员列表
		admin.GET("/admin_index.html", backend.GetAdminList)

		//添加管理员页面
		admin.GET("/admin_create.html", backend.AdminCreate)

		//添加管理员
		admin.POST("/add", backend.AdminAdd)

		//删除管理员
		admin.POST("/del", backend.AdminDel)

		//编辑管理员页面
		admin.GET("/admin_edit.html", backend.AdminEdit)

		//保存管理员信息
		admin.POST("/save", backend.AdminSave)

		//修改管理员密码页面
		admin.GET("/edit_password.html", backend.AdminEdit)

		//角色列表
		admin.GET("/role_index.html", backend.GetRoleList)

		//添加角色

		//路由列表
		admin.GET("/router_index.html", backend.GetRouterList)

		//添加路由
	}
}
