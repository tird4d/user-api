package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tird4d/user-api/repositories"
	"github.com/tird4d/user-api/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	userIDRaw, exists := c.Get("user_id")
	userID, ok := userIDRaw.(string)

	if !exists || !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid user ID"})
		c.Abort()
		return
	}

	repo := &repositories.MongoUserRepository{}
	user, err := services.GetUser(ctx, repo, userID)

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

func UpdateMeHandler(c *gin.Context) {
	//Creating a context
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	//checking user_id existence in request
	userIDRaw, exists := c.Get("user_id")
	userID, ok := userIDRaw.(string)

	if !exists || !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	//creating an objectId
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Cannot create object ID"})
		return
	}

	// Bind and validate input

	var body struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Calling the user service
	repo := &repositories.MongoUserRepository{}
	err = services.UpdateMe(ctx, repo, oid, body.Name, body.Email)

	//Sending response
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
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
