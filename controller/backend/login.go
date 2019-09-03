package backend

import (
	"GinApi/middleware/error"
	"GinApi/models"
	"GinApi/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

//后台登陆页
func AdminLoginIndex(c *gin.Context) {
	sessions := "ss"
	c.HTML(200, "login.html", gin.H{"sessions": sessions})
}

//登陆
func AdminLogin(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("passoword", "")
	code := error.INVALID_PARAMS

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	maps["user_name"] = username
	user, err := models.GetAdmin(maps)
	if err != nil {
		//用户不存在
		code := error.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  error.GetMsg(code),
			"data": data,
		})

		return
	}

	hashPassword := user.Password
	if !util.PasswordVerify(password, hashPassword) {
		//密码错误
		code := error.ERROR_NOT_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  error.GetMsg(code),
			"data": data,
		})
	}

	code = error.SUCCESS

	//生成token,session
	token, time, err := util.GenerateToken(user.Username, password)
	if err != nil {
		code = error.ERROR_AUTH_TOKEN
	} else {
		data["token"] = token
		data["exp_time"] = time
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  error.GetMsg(code),
		"data": data,
	})
}
