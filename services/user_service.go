package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/tird4d/user-api/models"
	"github.com/tird4d/user-api/repositories"
	"github.com/tird4d/user-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(repo repositories.UserRepository, name, email, password string) error {

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
	}

	_, err = repo.InsertNewUser(&user)

	return err
}

func LoginUser(repo repositories.UserRepository, email, password string) (string, error) {
	var user *models.User
	user, err := repo.FindByEmail(email)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		fmt.Println("token generated error:", err.Error())

		return "invalid token", err
	}

	return token, nil
}

func GetUser(user_id string) (models.User, error) {
	var user models.User

	oid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return user, err
	}

	err = models.UserCollection().FindOne(context.Background(), bson.M{"_id": oid}).Decode(&user)

	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, err

}

func GetAllUsers(ctx context.Context, repo repositories.UserRepository) ([]models.User, error) {
	return repo.GetAllUsers(ctx)
}
