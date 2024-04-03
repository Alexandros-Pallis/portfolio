package handler

import (
	"apallis/portfolio/model"
	"apallis/portfolio/renderer"
	"apallis/portfolio/view"
	"apallis/portfolio/view/pages"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{
    Permissions []string `form:"permissions[][]"`
}

func (h *UserHandler) Show(c *gin.Context) {
	var user model.User
	users, err := user.GetAll()
	if err != nil {
		log.Println("error getting users: ", err)
	}
	render := renderer.New(c, http.StatusOK, pages.ShowUserPage(users))
	c.Render(http.StatusOK, render)
	return
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
	render := renderer.New(c, http.StatusOK, view.Login(&user, flashes))
	c.Render(http.StatusOK, render)
	return
}

func (h *UserHandler) ManagePermissions(c *gin.Context) {
    var user model.User
    var permission model.Permission
    users, err := user.GetAll()
    if err != nil {
        log.Println("error getting users: ", err)
    }
    permissions, err := permission.GetAll()
    if err != nil {
        log.Println("error getting permissions: ", err)
    }
    render := renderer.New(c, http.StatusOK, pages.ManagePermissionsPage(users, permissions))
    c.Render(http.StatusOK, render)
}

func (h *UserHandler) AttemptManagePermissions(c *gin.Context) {
    // TODO: figure out how to bind the permissions of multidimensional post array
    // var permission model.Permission
    var foo any
    if err := c.ShouldBind(&foo); err != nil {
        log.Println("error binding permission: ", err)
    }
    log.Println("permissions: ", foo)
    c.JSON(http.StatusOK, gin.H{
        "message": "success",
        "permissions": foo,
    })
    // c.Redirect(http.StatusFound, "/manage-permissions")
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
