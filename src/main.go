package main

import (
	"apallis/portfolio/database"
	"apallis/portfolio/migration"
	"apallis/portfolio/model"
	"apallis/portfolio/renderer"
	"apallis/portfolio/routes"
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
	ginHTMLRenderer := router.HTMLRender
	router.HTMLRender = &renderer.HTMLTemplRenderer{
		FallbackHtmlRenderer: ginHTMLRenderer,
    }
	router.SetTrustedProxies(nil)
	router.Static("/dist", "./dist")
	db := database.Init()
	migration.Run(db)
	store := gormsessions.NewStore(db, true, []byte("secret"))
	router.Use(sessions.Sessions("session", store))
    gob.Register(model.User{})
	routes.Init(router)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":80")
}
