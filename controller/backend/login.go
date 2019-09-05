package backend

import (
	"GinApi/middleware/error"
	"GinApi/models"
	"GinApi/pkg/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

//后台登陆页
func AdminLoginIndex(c *gin.Context) {
	session := sessions.Default(c)
	uid := session.Get("user_id")
	c.HTML(200, "login.html", gin.H{"sessions": uid})
}

//登陆
func AdminLogin(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	validErr := validation.Errors{
		"username": validation.Validate(username, validation.Required.Error("username不能为空")),
		"password": validation.Validate(password, validation.Required.Error("password不能为空")),
	}.Filter()

	if validErr != nil {
		util.JsonErrResponse(c, error.INVALID_PARAMS)
		return
	}

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	maps["user_name"] = username

	user, err := models.GetAdmin(maps)
	if err != nil {
		//用户不存在
		util.JsonErrResponse(c, error.ERROR_NOT_EXIST_USER)
		return
	}

	hashPassword := user.Password
	if !util.PasswordVerify(password, hashPassword) {
		//密码错误
		util.JsonErrResponse(c, error.ERROR_NOT_EXIST_USER)
		return
	}

	if user.State == 2 || user.State == 3 {
		//用户被禁止/删除
		util.JsonErrResponse(c, error.ERROR_DISABLE_USER)
		return
	}

	//生成token
	token, time, err := util.GenerateToken(user.UserName, hashPassword)
	if err != nil {
		//Token错误
		util.JsonErrResponse(c, error.ERROR_AUTH_TOKEN)
		return
	} else {
		data["token"] = token
		data["exp_time"] = time
	}
	util.JsonSuccessResponse(c, error.SUCCESS, data)
}
