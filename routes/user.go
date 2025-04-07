package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tird4d/user-api/handlers"
	"github.com/tird4d/user-api/middlewares"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", handlers.RegisterHandler)

	r.POST("/login", handlers.LoginHandler)

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	auth.GET("/me", handlers.MeHandler)
	auth.PUT("/me", handlers.UpdateMeHandler)

	admin := auth.Group("/admin")
	admin.Use(middlewares.AuthorizeRole("admin"))
	admin.GET("/users", handlers.GetAllUsersHandler)

}
