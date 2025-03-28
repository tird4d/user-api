package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	err := services.RegisterUser(input.Name, input.Email, input.Password)
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
	token, err := services.LoginUser(input.Email, input.Password)
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
