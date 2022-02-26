package service

import "github.com/FakeqProject/fakeq-app/model"

type ImageServiceInterface interface {
	GetAllImage() ([]model.Image, error)
	GetAllImageWithPagination(offset, limit int) ([]model.Image, error)
	InsertNewImage(image model.Image) error
	GetImageById(id uint) (model.Image, error)
	UpdateImageById(id uint, values map[string]interface{}) error
	DeleteImageById(id uint) error
}
