package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tird4d/user-api/config"
	"github.com/tird4d/user-api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("⚠️ Error loading .env file")
	}

	config.ConnectDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "User API is running 🚀"})
	})

	routes.UserRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
