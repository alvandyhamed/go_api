package service

import (
	"errors"
	"go_api_learn/dto"
	"go_api_learn/model"
	"go_api_learn/repository"
	"go_api_learn/utils"
	"golang.org/x/crypto/bcrypt"
)

type authServiceImpl struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authServiceImpl{repo: repo}
}
func (s *authServiceImpl) Register(input dto.RegisterInput) (*model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  string(hashedPassword),
	}
	err = s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *authServiceImpl) Login(input dto.LoginInput) (string, error) {
	user, err := s.repo.GetByEmail(input.Email)
	if err != nil {
		return "", errors.New("ایمیل یا رمز نادرست است")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("ایمیل یا رمز نادرست است")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil

}
