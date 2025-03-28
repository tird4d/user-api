package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tird4d/user-api/handlers"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/register", handlers.RegisterHandler)
}

func LoginRoutes(r *gin.Engine) {
	r.POST("/login", handlers.LoginHandler)
}
