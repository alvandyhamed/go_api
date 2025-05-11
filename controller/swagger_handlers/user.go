package swagger_handlers

import (
	"github.com/gin-gonic/gin"
	"go_api_learn/controller"
	"go_api_learn/repository"
	"go_api_learn/service"
	"gorm.io/gorm"
)

type UserSwaggerHandler struct {
	DB *gorm.DB
}

// GetAll godoc
// @Summary دریافت لیست کاربران
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.User
// @Failure 401 {object} map[string]interface{}
// @Router /users [get]
func (h *UserSwaggerHandler) GetAll(c *gin.Context) {
	repo := repository.NewUserRepository(h.DB)
	userService := service.NewUserService(repo)
	userController := controller.NewUserController(userService)
	userController.GetAll(c)
}

// Search godoc
// @Summary جستجوی کاربران
// @Tags Users
// @Produce json
// @Param query query string true "کلمه کلیدی"
// @Security BearerAuth
// @Success 200 {array} model.User
// @Failure 401 {object} map[string]interface{}
// @Router /users/search [get]
func (h *UserSwaggerHandler) Search(c *gin.Context) {
	repo := repository.NewUserRepository(h.DB)
	userService := service.NewUserService(repo)
	userController := controller.NewUserController(userService)
	userController.Search(c)
}
