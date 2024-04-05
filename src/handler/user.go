package handler

import (
	"apallis/portfolio/model"
	"apallis/portfolio/renderer"
	"apallis/portfolio/view"
	"apallis/portfolio/view/pages"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (h *UserHandler) Add(c *gin.Context) {
	render := renderer.New(c, http.StatusOK, pages.AddUserPage())
	c.Render(http.StatusOK, render)
	return
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
	var user model.User
	contenxtUser, ok := c.Get("user")
	if !ok {
		user = model.User{}
	} else {
		user = contenxtUser.(model.User)
	}
	render := renderer.New(c, http.StatusOK, view.Login(&user))
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
	render := renderer.New(c, http.StatusOK, pages.ManagePermissionsPage(users, permissions, user.GetCurrentUser(c)))
	c.Render(http.StatusOK, render)
}

func (h *UserHandler) AttemptManagePermissions(c *gin.Context) {
	var user model.User
	session := sessions.Default(c)
	users, err := user.GetAll()
	if err != nil {
		session.AddFlash(model.NewFlash(err.Error(), "error"))
		session.Save()
		c.Redirect(http.StatusFound, "/users/permissions/manage")
		return
	}
	c.Request.ParseForm()
	form := c.Request.PostForm
	for _, user := range users {
		key := fmt.Sprintf("permissions[%d]", user.Id)
		permissions, ok := form[key]
		if !ok {
			continue
		}
        if err := user.RemovePermissions(); err != nil {
            session.AddFlash(model.NewFlash(err.Error(), "error"))
            session.Save()
            c.Redirect(http.StatusFound, "/users/permissions/manage")
            return
        }
		for _, permissionId := range permissions {
			var permission model.Permission
			if err := permission.GetById(permissionId); err != nil {
				session.AddFlash(model.NewFlash(err.Error(), "error"))
				continue
			}
			user.AddPermission(permission)
		}
		if err := user.Save(); err != nil {
			session.AddFlash(model.NewFlash(err.Error(), "error"))
			session.Save()
			c.Redirect(http.StatusFound, "/users/permissions/manage")
			return
		}
        if user.Id == user.GetCurrentUser(c).Id {
            session.Set("user", user)
        }
	}
	session.AddFlash(model.NewFlash("Permissions updated successfully", "success"))
	session.Save()
	c.Redirect(http.StatusFound, "/users/permissions/manage")
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
