package routes

import (
	"apallis/portfolio/handler"
	"apallis/portfolio/middleware"
	"apallis/portfolio/model"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	auth := middleware.AuthMiddleware{}
	// dashboard routes
	dashboardHandler := handler.DashboardHandler{}
	userHandler := handler.UserHandler{}
	authorized := router.Group("/dashboard")
	authorized.Use(auth.AuthRequired())
	{
		router.GET("dashboard", auth.AuthRequired(), dashboardHandler.Show)

		// user routes
		router.GET("users/add", auth.WithPermission(model.Write), userHandler.Add)
		router.GET("users/show", auth.WithPermission(model.Read), userHandler.Show)
		router.GET("users/permissions/manage", auth.WithPermission(model.Read), userHandler.ManagePermissions)
		router.POST("users/permissions/manage", auth.WithPermission(model.Delete), userHandler.AttemptManagePermissions)

        // project routes
        projectHandler := handler.ProjectHandler{}
        router.GET("projects/add", auth.WithPermission(model.Write), projectHandler.Add)
        router.POST("projects/add", auth.WithPermission(model.Write), projectHandler.AttempAdd)
	}

	// login routes
	router.GET("login", auth.AuthNotRequired(), userHandler.Login)
	router.GET("logout", auth.AuthRequired(), userHandler.Logout)
	router.POST("login", auth.AuthNotRequired(), userHandler.AttemptLogin)
}
