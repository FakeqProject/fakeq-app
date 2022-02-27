package service

import (
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/FakeqProject/fakeq-app/model"
)

func GetAllMessageImage() ([]model.MessageImage, error) {
	db := database.GetDB()
	var messageImages []model.MessageImage
	err := db.Find(&messageImages).Error

	return messageImages, err
}

func GetAllMessageImageWithPagination(offset, limit int) ([]model.MessageImage, error) {
	db := database.GetDB()
	var messageImages []model.MessageImage
	err := db.Offset(offset).Limit(limit).Find(&messageImages).Error

	return messageImages, err
}

func GetAllMessageImageByMessageID(messageID uint) ([]model.MessageImage, error) {
	db := database.GetDB()
	var messageImages []model.MessageImage
	err := db.Where("message_id = ?", messageID).Find(&messageImages).Error

	return messageImages, err
}

func InsertNewMessageImage(messageImage model.MessageImage) error {
	db := database.GetDB()
	err := db.Create(&messageImage).Error

	return err
}

func GetMessageImageByID(id uint) (model.MessageImage, error) {
	db := database.GetDB()
	var messageImage model.MessageImage
	err := db.First(&messageImage, id).Error

	return messageImage, err
}

func UpdateMessageImageByID(id uint, values map[string]interface{}) error {
	db := database.GetDB()
	err := db.Model(&model.MessageImage{}).Where("id = ?", id).Updates(values).Error

	return err
}

func DeleteMessageImageByID(id uint) error {
	db := database.GetDB()
	err := db.Delete(&model.MessageImage{}, id).Error

	return err
}

func DeleteMessageImageByMessageIDAndImageID(messageID, imageID uint) error {
	db := database.GetDB()
	err := db.Delete(&model.MessageImage{}, "message_id = ? AND image_id = ?", messageID, imageID).Error

	return err
}
