package controller

import (
	"go_api_learn/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در دریافت کاربران"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) Search(c *gin.Context) {
	query := c.Query("query")
	users, err := uc.userService.SearchUsers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در جستجو"})
		return
	}
	c.JSON(http.StatusOK, users)
}
