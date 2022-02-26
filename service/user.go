package service

import "github.com/FakeqProject/fakeq-app/model"

type UserServiceInterface interface {
	GetAllUser() ([]model.User, error)
	GetAllUserWithPagination(offset, limit int) ([]model.User, error)
	InsertNewUser(user model.User) error
	GetUserById(id uint) (model.User, error)
	UpdateUserById(id uint, values map[string]interface{}) error
	DeleteUserById(id uint) error
}
