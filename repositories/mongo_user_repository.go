package repositories

import (
	"context"
	"time"

	"github.com/tird4d/user-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct{}

func (r *MongoUserRepository) FindByID(ctx context.Context, oid primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := models.UserCollection().FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := models.UserCollection().FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) InsertNewUser(user *models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := models.UserCollection().InsertOne(ctx, user)

	defer cancel()

	return result, err
}

func (r *MongoUserRepository) UpdateUserByID(ctx context.Context, oid primitive.ObjectID, update bson.M) error {

	filter := bson.M{"_id": oid}
	_, err := models.UserCollection().UpdateOne(ctx, filter, bson.M{"$set": update})

	return err
}
func (r *MongoUserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {

	var users []models.User

	cursor, err := models.UserCollection().Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return users, nil

}
