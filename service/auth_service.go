package service

import (
	"go_api_learn/dto"
	"go_api_learn/model"
)

type AuthService interface {
	Register(input dto.RegisterInput) (*model.User, error)
	Login(input dto.LoginInput) (string, error)
}
