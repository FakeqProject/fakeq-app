package service

import (
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/FakeqProject/fakeq-app/model"
)

func GetAllUser() ([]model.User, error) {
	db := database.GetDB()
	var users []model.User
	err := db.Find(&users).Error

	return users, err
}

func GetAllUserWithPagination(offset, limit int) ([]model.User, error) {
	db := database.GetDB()
	var users []model.User
	err := db.Offset(offset).Limit(limit).Find(&users).Error

	return users, err
}

func InsertAllUser(user model.User) error {
	db := database.GetDB()
	err := db.Create(&user).Error

	return err
}

func GetUserByID(id uint) (model.User, error) {
	db := database.GetDB()
	var user model.User
	err := db.First(&user, id).Error

	return user, err
}

func UpdateUserByID(id uint, values map[string]interface{}) error {
	db := database.GetDB()
	err := db.Model(&model.User{}).Where("id = ?", id).Updates(values).Error

	return err
}

func DeleteUserByID(id uint) error {
	db := database.GetDB()
	err := db.Delete(&model.User{}, id).Error

	return err
}
