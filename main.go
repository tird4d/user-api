package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tird4d/user-api/config"
	"github.com/tird4d/user-api/routes"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User API is running 🚀"})
	})

	routes.UserRoutes(router)
	routes.LoginRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
