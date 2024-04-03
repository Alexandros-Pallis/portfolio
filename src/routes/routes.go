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
	router.GET("dashboard", dashboardHandler.Show)

	userHandler := handler.UserHandler{}
    // user routes
    router.GET("users/show", userHandler.Show)
    router.GET("users/permissions/manage", userHandler.ManagePermissions)
    router.POST("users/permissions/manage", userHandler.AttemptManagePermissions)

	// login routes
	router.GET("login", userHandler.Login)
	router.GET("logout", userHandler.Logout)
	router.POST("login", userHandler.AttemptLogin)
}
