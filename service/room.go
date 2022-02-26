package service

import "github.com/FakeqProject/fakeq-app/model"

type RoomServiceInterface interface {
	GetAllRoom() ([]model.Room, error)
	GetAllRoomWithPagination(offset, limit int) ([]model.Room, error)
	InsertNewRoom(room model.Room) error
	GetRoomById(id uint) (model.Room, error)
	UpdateRoomById(id uint, values map[string]interface{}) error
	DeleteRoomById(id uint) error
}
