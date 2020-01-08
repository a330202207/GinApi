package backend

import (
	"GinApi/middleware/casbin"
	"GinApi/model"
	"GinApi/package/error"
	"GinApi/service"
	"GinApi/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//后台登陆页
func AdminLoginIndex(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("user_id")
	c.HTML(200, "backend_login.html", gin.H{"user_id": userId})
}

//登陆
func AdminLogin(c *gin.Context) {
	var service service.AccountInfo
	if err := c.ShouldBind(&service); err != nil {
		util.JsonErrResponse(c, error.ERROR_NOT_EXIST_USER)
		return
	}

	admin, errCode := service.Login()
	if errCode != 200 {
		util.JsonErrResponse(c, errCode)
		return
	}

	nowLoginCnt := admin.LoginCnt
	loginInfo := model.Admin{
		LoginDate: time.Now(),
		LoginIp:   c.ClientIP(),
		LoginCnt:  nowLoginCnt + 1,
	}

	if err := model.UpdateLoginInfo(admin.ID, loginInfo); err != nil {
		util.JsonErrResponse(c, error.ERROR_SQL_UPDATE_FAIL)
	}

	s := sessions.Default(c)
	s.Clear()
	s.Set("admin_id", admin.ID)
	s.Save()

	err := casbin.AddRoleForUser(admin.ID)
	if err != nil {
		util.JsonErrResponse(c, error.ERROR)
	}

	util.JsonSuccessResponse(c, errCode, map[string]int{"admin_id": admin.ID})
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
