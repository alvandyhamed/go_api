package service

import (
	"go_api_learn/model"
	"go_api_learn/repository"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	SearchUsers(query string) ([]model.User, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) GetAllUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userServiceImpl) SearchUsers(query string) ([]model.User, error) {
	return s.repo.Search(query)
}
