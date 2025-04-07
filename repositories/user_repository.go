package repositories

import (
	"context"

	"github.com/tird4d/user-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	InsertNewUser(user *models.User) (*mongo.InsertOneResult, error)
	UpdateUserByID(ctx context.Context, oid primitive.ObjectID, update bson.M) error
	GetAllUsers(ctx context.Context) ([]models.User, error)
}
