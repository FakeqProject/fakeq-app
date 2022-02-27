package service

import (
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/FakeqProject/fakeq-app/model"
)

func GetAllMessageFile() ([]model.MessageFile, error) {
	db := database.GetDB()
	var messageFiles []model.MessageFile
	err := db.Find(&messageFiles).Error

	return messageFiles, err
}

func GetAllMessageFileWithPagination(offset, limit int) ([]model.MessageFile, error) {
	db := database.GetDB()
	var messageFiles []model.MessageFile
	err := db.Offset(offset).Limit(limit).Find(&messageFiles).Error

	return messageFiles, err
}

func GetAllMessageFileByMessageID(messageID uint) ([]model.MessageFile, error) {
	db := database.GetDB()
	var messageFiles []model.MessageFile
	err := db.Where("message_id = ?", messageID).Find(&messageFiles).Error

	return messageFiles, err
}

func InsertNewMessageFile(messageFile model.MessageFile) error {
	db := database.GetDB()
	err := db.Create(&messageFile).Error

	return err
}

func GetMessageFileByID(id uint) (model.MessageFile, error) {
	db := database.GetDB()
	var messageFile model.MessageFile
	err := db.First(&messageFile, id).Error

	return messageFile, err
}

func UpdateMessageFileByID(id uint, values map[string]interface{}) error {
	db := database.GetDB()
	err := db.Model(&model.MessageFile{}).Where("id = ?", id).Updates(values).Error

	return err
}

func DeleteMessageFileByID(id uint) error {
	db := database.GetDB()
	err := db.Delete(&model.MessageFile{}, id).Error

	return err
}

func DeleteMessageFileByMessageIDAndFileID(messageID, fileID uint) error {
	db := database.GetDB()
	err := db.Where("message_id = ? AND file_id = ?", messageID, fileID).Delete(&model.MessageFile{}).Error

	return err
}
