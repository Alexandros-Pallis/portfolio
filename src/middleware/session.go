package middleware

import (
	"apallis/portfolio/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type FlashesMiddleware struct{}

func (m *FlashesMiddleware) WithSessionData() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := model.User{}
		session := sessions.Default(c)
		sessionUser := session.Get("user")
		if sessionUser != nil {
			user = sessionUser.(model.User)
		}
		c.Set("user", user)
		var flashes []model.Flash
		sessionFlashes := session.Flashes()
		if sessionFlashes == nil || len(sessionFlashes) == 0 {
			c.Set("flashes", flashes)
		} else {
            for _, f := range sessionFlashes {
                flashes = append(flashes, f.(model.Flash))
            }
            c.Set("flashes", flashes)
        }
		session.Save()
		c.Next()
	}
}
