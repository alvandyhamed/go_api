package tests

import (
	"github.com/stretchr/testify/assert"
	"go_api_learn/model"
	"go_api_learn/repository/mocks"
	"go_api_learn/service"
	"testing"
)

func TestUserService_GetAllUsers(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	mockUsers := []model.User{
		{
			ID:        [16]byte{1}, // یا uuid.New() برای حالت واقعی
			FirstName: "Ali",
			LastName:  "Ebadi",
			Email:     "ali@example.com",
		},
		{
			ID:        [16]byte{2},
			FirstName: "Sara",
			LastName:  "Ahmadi",
			Email:     "sara@example.com",
		},
	}

	// وقتی GetAll صدا زده شد، این mockUsers رو برگردون
	mockRepo.On("GetAll").Return(mockUsers, nil)

	users, err := userService.GetAllUsers()

	assert.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, "Ali", users[0].FirstName)
	assert.Equal(t, "sara@example.com", users[1].Email)

	mockRepo.AssertExpectations(t)
}
func TestUserService_SearchUsers(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)

	query := "ali"
	mockUsers := []model.User{
		{
			ID:        [16]byte{1},
			FirstName: "Ali",
			LastName:  "Ebadi",
			Email:     "ali@example.com",
		},
	}

	// تنظیم رفتار mock
	mockRepo.On("Search", query).Return(mockUsers, nil)

	// اجرای تابع
	users, err := userService.SearchUsers(query)

	// بررسی نتیجه
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "ali@example.com", users[0].Email)

	mockRepo.AssertExpectations(t)
}
