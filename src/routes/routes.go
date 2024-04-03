package routes

import (
	"apallis/portfolio/handler"
	"apallis/portfolio/middleware"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	authMiddleware := middleware.AuthMiddleware{}
    router.Use(authMiddleware.AuthRequired())

	// dashboard routes
	dashboardHandler := handler.DashboardHandler{}
	router.GET("/dashboard", dashboardHandler.Show)

	// login routes
	userHandler := handler.UserHandler{}
	router.GET("/login", userHandler.Login)
	router.GET("/logout", userHandler.Logout)
	router.POST("/login", userHandler.AttemptLogin)
}
