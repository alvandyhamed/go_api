package tests

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go_api_learn/config"
	"go_api_learn/dto"
	"go_api_learn/model"
	"go_api_learn/repository/mocks"
	"go_api_learn/service"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestAuthService_Register(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)

	input := dto.RegisterInput{
		FirstName: "Ali",
		LastName:  "Ebadi",
		Email:     "ali@example.com",
		Password:  "123456",
	}

	mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return(nil)

	user, err := authService.Register(input)

	assert.NoError(t, err)
	assert.Equal(t, "Ali", user.FirstName)
	assert.Equal(t, "Ebadi", user.LastName)
	assert.Equal(t, "ali@example.com", user.Email)
	assert.NotEmpty(t, user.Password) // هش شده
	mockRepo.AssertExpectations(t)
}

func TestAuthService_Login_Success(t *testing.T) {
	// ✅ پیکربندی تست
	config.AppConfig = &config.Config{
		JwtSecret: "test_secret",
	}

	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	user := &model.User{
		ID:        uuid.New(),
		Email:     "ali@example.com",
		Password:  string(hashedPassword),
		FirstName: "Ali",
		LastName:  "Ebadi",
	}

	input := dto.LoginInput{
		Email:    "ali@example.com",
		Password: "123456",
	}

	mockRepo.On("GetByEmail", "ali@example.com").Return(user, nil)

	token, err := authService.Login(input)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}
