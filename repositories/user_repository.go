package repositories

import (
	"github.com/tird4d/user-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	InsertNewUser(user *models.User) (*mongo.InsertOneResult, error)
}
