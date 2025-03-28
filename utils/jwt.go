package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateJWT(userId primitive.ObjectID) (string, error) {
	_ = userId

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		secret = "default_secret" //only for development
	}

	claims := jwt.MapClaims{
		"user_id": userId.Hex(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	fmt.Println(token)
	return tokenString, err

}
