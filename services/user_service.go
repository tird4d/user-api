package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/tird4d/user-api/models"
	"github.com/tird4d/user-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(name, email, password string) error {

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	_, err = models.UserCollection().InsertOne(ctx, user)

	defer cancel()

	return err
}

func LoginUser(email, password string) (string, error) {
	var user models.User
	err := models.UserCollection().FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		fmt.Println("token generated error:", err.Error())

		return "invalid token", err
	}

	return token, nil
}
