package service

import (
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/FakeqProject/fakeq-app/model"
)

func GetAllImage() ([]model.Image, error) {
	db := database.GetDB()
	var images []model.Image
	err := db.Find(&images).Error

	return images, err
}

func GetAllImageWithPagination(offset, limit int) ([]model.Image, error) {
	db := database.GetDB()
	var images []model.Image
	err := db.Offset(offset).Limit(limit).Find(&images).Error

	return images, err
}

func InsertNewImage(image model.Image) error {
	db := database.GetDB()
	err := db.Create(&image).Error

	return err
}

func GetImageByID(id uint) (model.Image, error) {
	db := database.GetDB()
	var image model.Image
	err := db.First(&image, id).Error

	return image, err
}

func UpdateImageByID(id uint, values map[string]interface{}) error {
	db := database.GetDB()
	err := db.Model(&model.Image{}).Where("id = ?", id).Updates(values).Error

	return err
}

func DeleteImageByID(id uint) error {
	db := database.GetDB()
	err := db.Delete(&model.Image{}, id).Error

	return err
}
