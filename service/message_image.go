package service

import "github.com/FakeqProject/fakeq-app/model"

type MessageImageServiceInterface interface {
	GetAllMessageImage() ([]model.MessageImage, error)
	GetAllMessageImageWithPagination(offset, limit int) ([]model.MessageImage, error)
	InsertNewMessageImage(messageImage model.MessageImage) error
	GetMessageImageById(id uint) (model.MessageImage, error)
	UpdateMessageImageById(id uint, values map[string]interface{}) error
	DeleteMessageImageById(id uint) error
}
