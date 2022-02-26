package service

import "github.com/FakeqProject/fakeq-app/model"

type MessageFileServiceInterface interface {
	GetAllMessageFile() ([]model.MessageFile, error)
	GetAllMessageFileWithPagination(offset, limit int) ([]model.MessageFile, error)
	InsertNewMessageFile(messageFile model.MessageFile) error
	GetMessageFileById(id uint) (model.MessageFile, error)
	UpdateMessageFileById(id uint, values map[string]interface{}) error
	DeleteMessageFileById(id uint) error
}
