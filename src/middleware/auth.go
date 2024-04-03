package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func (m *AuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		authenticated := session.Get("authenticated")
		if authenticated == nil && c.Request.URL.Path != "/login"{
            session.AddFlash("You need to login first")
            session.Save()
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
        if authenticated != nil && c.Request.URL.Path == "/login" {
            c.Redirect(http.StatusFound, "/dashboard")
            c.Abort()
            return
        }
        c.Set("user", session.Get("user"))
	}
}
