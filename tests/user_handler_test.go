package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_api_learn/config"
	"go_api_learn/controller"
	"go_api_learn/middleware"
	"go_api_learn/model"
	"go_api_learn/repository/mocks"
	"go_api_learn/service"
	"go_api_learn/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserSearchHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.AppConfig = &config.Config{JwtSecret: "test_secret"}
	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)
	userController := controller.NewUserController(userService)

	// داده‌ی برگشتی از repo
	mockUsers := []model.User{
		{FirstName: "Ali", Email: "ali@example.com"},
		{FirstName: "Alicia", Email: "alicia@example.com"},
	}

	mockRepo.On("Search", "ali").Return(mockUsers, nil)

	router := gin.Default()

	// middleware JWT fake
	router.Use(func(c *gin.Context) {
		c.Set("user_id", "123") // شبیه‌سازی user از توکن
		c.Next()
	})

	router.GET("/users/search", userController.Search)

	req, _ := http.NewRequest(http.MethodGet, "/users/search?query=ali", nil)
	req.Header.Set("Authorization", "Bearer faketoken")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}
func TestUserGetAllHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	config.AppConfig = &config.Config{JwtSecret: "test_secret"}

	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)
	userController := controller.NewUserController(userService)

	// داده ساختگی
	mockUsers := []model.User{
		{FirstName: "Ali", Email: "ali@example.com"},
		{FirstName: "Sara", Email: "sara@example.com"},
	}

	mockRepo.On("GetAll").Return(mockUsers, nil)

	router := gin.Default()

	// شبیه‌سازی middleware که user_id رو داخل context می‌ذاره
	router.Use(func(c *gin.Context) {
		c.Set("user_id", "123")
		c.Next()
	})

	router.GET("/users", userController.GetAll)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Authorization", "Bearer faketoken")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}
func TestUserGetAll_WithRealJWTMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// ست کردن JWT Secret
	config.AppConfig = &config.Config{JwtSecret: "test_secret"}

	mockRepo := new(mocks.MockUserRepository)
	userService := service.NewUserService(mockRepo)
	userController := controller.NewUserController(userService)

	// کاربران فرضی
	mockUsers := []model.User{
		{FirstName: "Ali", Email: "ali@example.com"},
		{FirstName: "Sara", Email: "sara@example.com"},
	}

	mockRepo.On("GetAll").Return(mockUsers, nil)

	// تولید JWT واقعی
	token, err := utils.GenerateToken("123")
	assert.NoError(t, err)

	// ساخت router و افزودن middleware اصلی
	router := gin.Default()
	router.GET("/users", middleware.JWTAuthMiddleware(), userController.GetAll)

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockRepo.AssertExpectations(t)
}
