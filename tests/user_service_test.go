package tests

import (
	"errors"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tird4d/user-api/mocks"
	"github.com/tird4d/user-api/models"
	"github.com/tird4d/user-api/services"
	"github.com/tird4d/user-api/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("❌ Failed to load .env file")
	}
	// config.ConnectDB()

}

//Real Database test
/*
func TestUserLogin(t *testing.T) {

	repo := &repositories.MongoUserRepository{}
	_ = repo

	email := "ali@example.com"
	password := "supersecret"
	token, err := LoginUser(repo, email, password)

	if err != nil {
		t.Error("error", err.Error())
	}

	if token == "" {
		t.Error("error", "login failed")
	}

}

*/

//Test with no mocking package
/*
func TestLoginUser_WithMock_Success(t *testing.T) {
	password := "123456"
	hashed, _ := utils.HashPassword(password)

	mockRepo := &mocks.MockUserRepository{
		FakeUsers: map[string]*models.User{
			"ali@example.com": {
				Name:     "Ali",
				Email:    "ali@example.com",
				Password: hashed,
			},
		},
	}

	token, err := LoginUser(mockRepo, "ali@example.com", password)
	if err != nil {
		t.Fatalf("Expected login to succeed, got error: %v", err)
	}

	if token == "" {
		t.Error("Expected token, got empty string")
	}
}

func TestLoginUser_WithMock_WrongPassword(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		FakeUsers: map[string]*models.User{
			"ali@example.com": {
				Name:     "Ali",
				Email:    "ali@example.com",
				Password: "$2a$14$somethingWrongHashed", //  fake
			},
		},
	}

	_, err := LoginUser(mockRepo, "ali@example.com", "wrongPass")
	if err == nil {
		t.Error("Expected error for wrong password, got nil")
	}
}

func TestLoginUser_WithMock_UserNotFound(t *testing.T) {
	mockRepo := &mocks.MockUserRepository{
		FakeUsers: map[string]*models.User{},
	}

	_, err := LoginUser(mockRepo, "notfound@example.com", "any")
	if err == nil {
		t.Error("Expected error for non-existing user, got nil")
	}
}
*/

func TestRegisterUser_WithTestifyMock_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	user := models.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "123456",
	}

	mockRepo.On("InsertNewUser", mock.MatchedBy(func(u *models.User) bool {
		return u.Email == "test@test.com" && u.Name == "test"
	})).Return(&mongo.InsertOneResult{}, nil)

	//mockRepo.On("InsertNewUser", mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	err := services.RegisterUser(mockRepo, user.Name, user.Email, user.Password)

	assert.NoError(t, err)

}

func TestLoginUser_WithTestifyMock_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	password := "123456"
	hashed, _ := utils.HashPassword(password)

	mockRepo.On("FindByEmail", "ali@example.com").Return(&models.User{
		Email:    "ali@example.com",
		Password: hashed,
	}, nil)

	token, err := services.LoginUser(mockRepo, "ali@example.com", password)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	mockRepo.AssertExpectations(t)
}

func TestLoginUser_WithTestifyMock_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	mockRepo.On("FindByEmail", "notfound@example.com").Return(nil, errors.New("user not found"))

	_, err := services.LoginUser(mockRepo, "notfound@example.com", "any")
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}
