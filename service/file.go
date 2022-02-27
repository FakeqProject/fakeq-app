package service

import (
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/FakeqProject/fakeq-app/model"
)

func GetAllFile() ([]model.File, error) {
	db := database.GetDB()
	var files []model.File
	err := db.Find(&files).Error

	return files, err
}

func GetAllFileWithPagination(offset, limit int) ([]model.File, error) {
	db := database.GetDB()
	var files []model.File
	err := db.Offset(offset).Limit(limit).Find(&files).Error

	return files, err
}

func InsertNewFile(file model.File) error {
	db := database.GetDB()
	err := db.Create(&file).Error

	return err
}

func GetFileById(id uint) (model.File, error) {
	db := database.GetDB()
	var file model.File
	err := db.First(&file, id).Error

	return file, err
}

func UpdateFileById(id uint, values map[string]interface{}) error {
	db := database.GetDB()
	err := db.Model(&model.File{}).Where("id = ?", id).Updates(values).Error

	return err
}

func DeleteFileById(id uint) error {
	db := database.GetDB()
	err := db.Delete(&model.File{}, id).Error

	return err
}
