package repository

import (
	"go_api_learn/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]model.User, error)
	Search(query string) ([]model.User, error)
}
