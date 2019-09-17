package casbin

import (
	//"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//"github.com/casbin/casbin"
	"net/http"
	"strings"
)

func CheckLoginHandle() gin.HandlerFunc {
	return func(c *gin.Context) {

		//e := casbin.NewEnforcer("./configs/rbac.conf", "./configs/rbac_policy.csv")

		//获取请求的URL
		objUrl := c.Request.URL.RequestURI()

		if strings.HasPrefix(objUrl, "/static") || strings.HasPrefix(objUrl, "/admin/backend_login.html") || strings.HasPrefix(objUrl, "/admin/login") {
			c.Next()
			return
		}

		//获取请求方法
		//act := c.Request.Method

		//获取用户角色
		s := sessions.Default(c)

		userId := s.Get("user_id")

		if userId == nil {
			c.Redirect(http.StatusTemporaryRedirect, "/admin/backend_login.html")
			return
		}

		//管理员角色
		//sub := userId.(string)
		//
		////判断权限
		//if e.Enforce(sub, objUrl, act) {
		//	fmt.Println("通过权限")
		//	c.Next()
		//} else {
		//	fmt.Println("权限没有通过")
		//	c.Abort()
		//}
	}
}
