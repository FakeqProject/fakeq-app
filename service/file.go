package service

import "github.com/FakeqProject/fakeq-app/model"

type FileServiceInterface interface {
	GetAllFile() ([]model.File, error)
	GetAllFileWithPagination(offset, limit int) ([]model.File, error)
	InsertNewFile(file model.File) error
	GetFileById(id uint) (model.File, error)
	UpdateFileById(id uint, values map[string]interface{}) error
	DeleteFileById(id uint) error
}
