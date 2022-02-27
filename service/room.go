package service

import (
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/FakeqProject/fakeq-app/model"
)

func GetAllRoom() ([]model.Room, error) {
	db := database.GetDB()
	var rooms []model.Room
	err := db.Find(&rooms).Error

	return rooms, err
}

func GetAllRoomWithPagination(offset, limit int) ([]model.Room, error) {
	db := database.GetDB()
	var rooms []model.Room
	err := db.Offset(offset).Limit(limit).Find(&rooms).Error

	return rooms, err
}

func InsertNewRoom(room model.Room) error {
	db := database.GetDB()
	err := db.Create(&room).Error

	return err
}

func GetRoomById(id uint) (model.Room, error) {
	db := database.GetDB()
	var room model.Room
	err := db.First(&room, id).Error

	return room, err
}

func UpdateRoomById(id uint, values map[string]interface{}) error {
	db := database.GetDB()
	err := db.Model(&model.Room{}).Where("id = ?", id).Updates(values).Error

	return err
}

func DeleteRoomById(id uint) error {
	db := database.GetDB()
	err := db.Delete(&model.Room{}, id).Error

	return err
}
