package controller

import (
	"github.com/gin-gonic/gin"
	"go_api_learn/dto"
	"go_api_learn/service"
	"net/http"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر است", "details": err.Error()})
		return
	}
	user, err := ac.authService.Register(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "خطا در ثبت‌نام", "details": err.Error()})
		return

	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "ثبت‌نام با موفقیت انجام شد",
		"user": gin.H{
			"id":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
		},
	})

}

func (ac *AuthController) Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ورودی نامعتبر است", "details": err.Error()})
		return

	}
	token, err := ac.authService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ایمیل یا رمز اشتباه است"})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ورود موفقیت‌آمیز",
		"token":   token,
	})

}
