package session

import (
	"GinApi/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// Session
func Session() gin.HandlerFunc {
	store := cookie.NewStore([]byte(config.ServerSetting.SessionStore))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 7 * 86400, Path: "/"})
	return sessions.Sessions(config.ServerSetting.SessionName, store)
}
