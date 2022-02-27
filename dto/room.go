package dto

import "github.com/FakeqProject/fakeq-app/model"

type RoomResponseDto struct {
	ID     uint   `json:"id"`
	RoomID uint   `json:"room_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

type RoomRequestDto struct {
	RoomID uint   `json:"room_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
}

func CreateRoomResponseDto(room model.Room) RoomResponseDto {
	return RoomResponseDto{
		ID:     room.ID,
		RoomID: room.RoomID,
		Name:   room.Name,
		Type:   room.Type,
	}
}
