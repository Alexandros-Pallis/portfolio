package middleware

import (
	"apallis/portfolio/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct{}

func (m *AuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		authenticated := session.Get("authenticated")
		if authenticated == nil {
            session.AddFlash(model.NewFlash("You must be logged in to access this page", model.Error))
            session.Save()
            c.Redirect(http.StatusFound, "/login")
            c.Abort()
        }
		c.Next()
	}
}

func (m *AuthMiddleware) AuthNotRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        authenticated := session.Get("authenticated")
        if authenticated != nil {
            session.AddFlash(model.NewFlash("You are already logged in", model.Info))
            session.Save()
            c.Redirect(http.StatusFound, "/dashboard")
            c.Abort()
        }
        c.Next()
    }
}

func (m *AuthMiddleware) WithPermission(permissionType model.PermissionType) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user model.User
        current := user.GetCurrentUser(c)
        if !current.HasPermission(permissionType) {
            session := sessions.Default(c)
            session.AddFlash(model.NewFlash("You do not have permission to access this page", model.Error))
            session.Save()
            c.Redirect(http.StatusFound, "/dashboard")
            c.Abort()
            return
        }
        c.Next()
    }
}
