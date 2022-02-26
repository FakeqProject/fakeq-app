package service

import "github.com/FakeqProject/fakeq-app/model"

type MessageServiceInterface interface {
	GetAllMessage() ([]model.Message, error)
	GetAllMessageWithPagination(offset, limit int) ([]model.Message, error)
	InsertNewMessage(message model.Message) error
	GetMessageById(id uint) (model.Message, error)
	UpdateMessageById(id uint, values map[string]interface{}) error
	DeleteMessageById(id uint) error
}
