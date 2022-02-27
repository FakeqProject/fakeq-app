package service

import (
	"github.com/FakeqProject/fakeq-app/database"
	"github.com/FakeqProject/fakeq-app/model"
)

func GetAllGroup() ([]model.Group, error) {
	db := database.GetDB()
	var groups []model.Group
	err := db.Find(&groups).Error

	return groups, err
}

func GetAllGroupWithPagination(offset, limit int) ([]model.Group, error) {
	db := database.GetDB()
	var groups []model.Group
	err := db.Offset(offset).Limit(limit).Find(&groups).Error

	return groups, err
}

func InsertNewGroup(group model.Group) error {
	db := database.GetDB()
	err := db.Create(&group).Error

	return err
}

func GetGroupByID(id uint) (model.Group, error) {
	db := database.GetDB()
	var group model.Group
	err := db.First(&group, id).Error

	return group, err
}

func UpdateGroupByID(id uint, values map[string]interface{}) error {
	db := database.GetDB()
	err := db.Model(&model.Group{}).Where("id = ?", id).Updates(values).Error

	return err
}

func DeleteGroupByID(id uint) error {
	db := database.GetDB()
	err := db.Delete(&model.Group{}, id).Error

	return err
}
