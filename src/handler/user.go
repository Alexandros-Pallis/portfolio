package handler

import (
	"apallis/portfolio/model"
	"apallis/portfolio/renderer"
	"apallis/portfolio/view"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (h *UserHandler) AuthRequiredMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		authenticated := session.Get("authenticated")
		if authenticated == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	user := model.User{}
	session := sessions.Default(c)
	flashes := session.Flashes()
	sessionUser := session.Get("user")
	session.Save()
	if sessionUser != nil {
		user = sessionUser.(model.User)
	}
	renderer := renderer.New(c, http.StatusOK, view.Login(&user, flashes))
	c.Render(http.StatusOK, renderer)
	return
}

func (h *UserHandler) AttemptLogin(c *gin.Context) {
	var user model.User
	session := sessions.Default(c)
	if err := c.ShouldBind(&user); err != nil {
		session.AddFlash(err.Error())
		session.Save()
		c.Redirect(http.StatusFound, "/login")
		return
	}
	session.Set("user", user)
	user, err := model.GetUserByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		session.AddFlash(err.Error())
		if err := session.Save(); err != nil {
			log.Println("error saving session: ", err)
		}
		c.Redirect(http.StatusFound, "/login")
		return
	}
	session.Set("user", user)
	session.Set("authenticated", true)
	err = session.Save()
	if err != nil {
		log.Println("error saving session: ", err)
	}
	c.Redirect(http.StatusFound, "/dashboard")
}

func (h *UserHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.Redirect(http.StatusFound, "/login")
}
