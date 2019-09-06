package session

import (
	"GinApi/pkg/setting"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session
func Session() gin.HandlerFunc {
	store := cookie.NewStore([]byte(setting.ServerSetting.SessionStore))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions(setting.ServerSetting.SessionName, store)
}
