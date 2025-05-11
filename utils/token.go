package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"go_api_learn/config"
	"time"
)

func GenerateToken(userID interface{}) (string, error) {
	secret := config.AppConfig.JwtSecret
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}
