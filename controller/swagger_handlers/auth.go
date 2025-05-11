package swagger_handlers

import (
	"github.com/gin-gonic/gin"
	"go_api_learn/controller"
	"go_api_learn/repository"
	"go_api_learn/service"
	"gorm.io/gorm"
)

type AuthSwaggerHandler struct {
	DB *gorm.DB
}

// Register godoc
// @Summary ثبت‌نام Swagger
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.RegisterInput true "اطلاعات ثبت‌نام"
// @Success 201 {object} map[string]interface{}
// @Failure 400,500 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *AuthSwaggerHandler) Register(c *gin.Context) {
	userRepo := repository.NewUserRepository(h.DB)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	authController.Register(c)
}

// Login godoc
// @Summary ورود کاربر
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body dto.LoginInput true "ورودی ورود"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *AuthSwaggerHandler) Login(c *gin.Context) {
	userRepo := repository.NewUserRepository(h.DB)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	authController.Login(c)
}
