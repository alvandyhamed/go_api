package tests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go_api_learn/config"
	"go_api_learn/controller"
	"go_api_learn/dto"
	"go_api_learn/model"
	"go_api_learn/repository/mocks"
	"go_api_learn/service"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)
	authController := controller.NewAuthController(authService)

	// داده‌ی ورودی
	input := dto.RegisterInput{
		FirstName: "Ali",
		LastName:  "Ebadi",
		Email:     "ali@example.com",
		Password:  "123456",
	}
	jsonInput, _ := json.Marshal(input)

	// شبیه‌سازی ایجاد موفق کاربر
	mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return(nil)

	// ساخت router و تعریف مسیر
	router := gin.Default()
	router.POST("/auth/register", authController.Register)

	// اجرای تست
	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)
	mockRepo.AssertExpectations(t)
}
func TestRegisterHandler_EmptyPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)
	authController := controller.NewAuthController(authService)

	input := dto.RegisterInput{
		FirstName: "Ali",
		LastName:  "Ebadi",
		Email:     "ali@example.com",
		Password:  "", // ⛔ خالی
	}
	jsonInput, _ := json.Marshal(input)

	router := gin.Default()
	router.POST("/auth/register", authController.Register)

	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
func TestRegisterHandler_RepoError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)
	authController := controller.NewAuthController(authService)

	input := dto.RegisterInput{
		FirstName: "Ali",
		LastName:  "Ebadi",
		Email:     "ali@example.com",
		Password:  "123456",
	}
	jsonInput, _ := json.Marshal(input)

	// ❌ Repo خطا می‌ده (مثلاً ایمیل تکراری)
	mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return(assert.AnError)

	router := gin.Default()
	router.POST("/auth/register", authController.Register)

	req, _ := http.NewRequest(http.MethodPost, "/auth/register", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
	mockRepo.AssertExpectations(t)
}
func TestLoginHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// تنظیم JWT برای تست
	config.AppConfig = &config.Config{
		JwtSecret: "test_secret",
	}

	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)
	authController := controller.NewAuthController(authService)

	// ساخت پسورد هش‌شده واقعی
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	user := &model.User{
		ID:       [16]byte{1},
		Email:    "ali@example.com",
		Password: string(hashedPassword),
	}

	// ورودی درخواست
	input := dto.LoginInput{
		Email:    "ali@example.com",
		Password: "123456",
	}
	jsonInput, _ := json.Marshal(input)

	// شبیه‌سازی repo
	mockRepo.On("GetByEmail", "ali@example.com").Return(user, nil)

	router := gin.Default()
	router.POST("/auth/login", authController.Login)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}
func TestLoginHandler_UserNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)
	authController := controller.NewAuthController(authService)

	input := dto.LoginInput{
		Email:    "notfound@example.com",
		Password: "123456",
	}
	jsonInput, _ := json.Marshal(input)

	mockRepo.On("GetByEmail", input.Email).Return((*model.User)(nil), assert.AnError)

	router := gin.Default()
	router.POST("/auth/login", authController.Login)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
}
func TestLoginHandler_WrongPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(mocks.MockUserRepository)
	authService := service.NewAuthService(mockRepo)
	authController := controller.NewAuthController(authService)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpass"), bcrypt.DefaultCost)
	user := &model.User{
		ID:       [16]byte{1},
		Email:    "ali@example.com",
		Password: string(hashedPassword),
	}

	input := dto.LoginInput{
		Email:    "ali@example.com",
		Password: "wrongpass",
	}
	jsonInput, _ := json.Marshal(input)

	mockRepo.On("GetByEmail", input.Email).Return(user, nil)

	router := gin.Default()
	router.POST("/auth/login", authController.Login)

	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusUnauthorized, resp.Code)
	mockRepo.AssertExpectations(t)
}
