package backend

import (
	"GinApi/middleware/error"
	"GinApi/pkg/util"
	"GinApi/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//后台登陆页
func AdminLoginIndex(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	c.HTML(200, "login.html", gin.H{"user_id": userId})
}

//登陆
func AdminLogin(c *gin.Context) {
	var service service.AdminLoginService
	if err := c.ShouldBind(&service); err == nil {
		if user, errCode := service.Login(); errCode != 200 {
			util.JsonErrResponse(c, errCode)
		} else {
			//登陆成功
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.ID)
			s.Save()
			util.JsonSuccessResponse(c, errCode, map[string]int{"user_id": user.ID})
		}
	} else {
		util.JsonErrResponse(c, error.ERROR_NOT_EXIST_USER)
	}
}
