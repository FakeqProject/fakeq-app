package service

import (
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/FakeqProject/fakeq-app/model"
)

func GetAllMessage() ([]model.Message, error) {
	db := database.GetDB()
	var messages []model.Message
	err := db.Find(&messages).Error

	return messages, err
}

func GetAllMessageWithPagination(offset, limit int) ([]model.Message, error) {
	db := database.GetDB()
	var messages []model.Message
	err := db.Offset(offset).Limit(limit).Find(&messages).Error

	return messages, err
}

func InsertNewMessage(message model.Message) error {
	db := database.GetDB()
	err := db.Create(&message).Error

	return err
}

func GetMessageByID(id uint) (model.Message, error) {
	db := database.GetDB()
	var message model.Message
	err := db.First(&message, id).Error

	return message, err
}

func UpdateMessageByID(id uint, values map[string]interface{}) error {
	db := database.GetDB()
	err := db.Model(&model.Message{}).Where("id = ?", id).Updates(values).Error

	return err
}

func DeleteMessageByID(id uint) error {
	db := database.GetDB()
	err := db.Delete(&model.Message{}, id).Error

	return err
}
