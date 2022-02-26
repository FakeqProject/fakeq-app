package service

import "github.com/FakeqProject/fakeq-app/model"

type GroupServiceInterface interface {
	GetAllGroup() ([]model.Group, error)
	GetAllGroupWithPagination(offset, limit int) ([]model.Group, error)
	InsertNewGroup(model.Group) error
	GetGroupById(id uint) (model.Group, error)
	UpdateGroupById(id uint, values map[string]interface{}) error
	DeleteGroupById(id uint) error
}
