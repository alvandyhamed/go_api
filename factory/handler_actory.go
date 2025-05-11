package factory

import (
	"github.com/gin-gonic/gin"
	"go_api_learn/controller"
	"go_api_learn/repository"
	"go_api_learn/service"
	"gorm.io/gorm"
)

type HandlerFactory struct {
	DB *gorm.DB
}

// 📌 این تابع کنترلر Auth رو می‌سازه و هندلرش رو برمی‌گردونه
func (hf *HandlerFactory) AuthRegisterHandler() func(c *gin.Context) {
	userRepo := repository.NewUserRepository(hf.DB)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	return authController.Register
}
