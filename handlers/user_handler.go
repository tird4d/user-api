package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tird4d/user-api/repositories"
	"github.com/tird4d/user-api/services"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func RegisterHandler(c *gin.Context) {

	var input RegisterInput

	// Bind and validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Send input to service layer
	repo := &repositories.MongoUserRepository{}
	err := services.RegisterUser(repo, input.Name, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// send response
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})

}

func LoginHandler(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Send input to service layer
	repo := &repositories.MongoUserRepository{}
	token, err := services.LoginUser(repo, input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "The user or password is invalid"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"token":   token,
			"message": "Login successful",
		})

	}

}

func MeHandler(c *gin.Context) {
	userIDRaw, exists := c.Get("user_id")
	userID, ok := userIDRaw.(string)

	if !exists || !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
		c.Abort()
		return
	}

	user, err := services.GetUser(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user.Name,
		"message": "this is user profile",
	})
}

func GetAllUsersHandler(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	//Send input to service layer
	repo := &repositories.MongoUserRepository{}
	users, err := services.GetAllUsers(ctx, repo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})

}
