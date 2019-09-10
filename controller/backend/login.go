package backend

import (
	"GinApi/package/error"
	"GinApi/service"
	"GinApi/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

//后台登陆页
func AdminLoginIndex(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	c.HTML(200, "login.html", gin.H{"user_id": userId})
}

//登陆
func AdminLogin(c *gin.Context) {
	var service service.AdminLoginInfo
	if err := c.ShouldBind(&service); err == nil {
		if admin, errCode := service.Login(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			//登陆成功
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", admin.ID)
			s.Save()
			util.JsonSuccessResponse(c, errCode, map[string]int{"user_id": admin.ID})
		}
	} else {
		util.JsonErrResponse(c, error.ERROR_NOT_EXIST_USER)
	}
}

//登出
func AdminLogOut(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": error.SUCCESS,
		"msg":  error.GetMsg(error.SUCCESS),
	})
}
