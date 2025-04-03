package mocks

//Without package mock handling
/*
import (
	"errors"

	"github.com/tird4d/user-api/models"
)

type MockUserRepository struct {
	FakeUsers map[string]*models.User
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	user, exists := m.FakeUsers[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
*/

import (
	"context"
	"github.com/stretchr/testify/mock"
	"github.com/tird4d/user-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if user, ok := args.Get(0).(*models.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) InsertNewUser(user *models.User) (*mongo.InsertOneResult, error) {
	args := m.Called(user)

	if result, ok := args.Get(0).(*mongo.InsertOneResult); ok {
		return result, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UserRepositoryMock) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return nil, nil
}
