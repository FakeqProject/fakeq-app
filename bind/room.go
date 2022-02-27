package bind

import (
	"github.com/FakeqProject/fakeq-app/dto"
	"github.com/FakeqProject/fakeq-app/model"
	"github.com/FakeqProject/fakeq-app/service"
	"github.com/webview/webview"
)

func RoomBindCollection(w webview.WebView) {
	w.Bind("getAllRoom", GetAllRoom())
	w.Bind("getAllRoomWithPagination", GetAllRoomWithPagination())
	w.Bind("getRoomByID", GetRoomByID())
	w.Bind("insertNewRoom", InsertNewRoom())
	w.Bind("updateRoomByID", UpdateRoomByID())
	w.Bind("deleteRoomByID", DeleteRoomByID())
}

func GetAllRoom() interface{} {
	return func() ([]dto.RoomResponseDto, error) {
		rooms, err := service.GetAllRoom()
		if err != nil {
			return nil, err
		}
		var roomDtos []dto.RoomResponseDto
		for _, room := range rooms {
			roomDtos = append(roomDtos, dto.CreateRoomResponseDto(room))
		}
		return roomDtos, nil
	}
}

func GetAllRoomWithPagination() interface{} {
	return func(page, limit int) ([]dto.RoomResponseDto, error) {
		rooms, err := service.GetAllRoomWithPagination((page-1)*limit, limit)
		if err != nil {
			return nil, err
		}
		var roomDtos []dto.RoomResponseDto
		for _, room := range rooms {
			roomDtos = append(roomDtos, dto.CreateRoomResponseDto(room))
		}
		return roomDtos, nil
	}
}

func GetRoomByID() interface{} {
	return func(id uint) (dto.RoomResponseDto, error) {
		room, err := service.GetRoomByID(id)
		if err != nil {
			return dto.RoomResponseDto{}, err
		}
		return dto.CreateRoomResponseDto(room), nil
	}
}

func InsertNewRoom() interface{} {
	return func(request dto.RoomRequestDto) (dto.RoomResponseDto, error) {
		room := model.Room{
			RoomID: request.RoomID,
			Name:   request.Name,
			Type:   request.Type,
		}
		err := service.InsertNewRoom(room)
		if err != nil {
			return dto.RoomResponseDto{}, err
		}
		return dto.CreateRoomResponseDto(room), nil
	}
}

func UpdateRoomByID() interface{} {
	return func(id uint, values map[string]interface{}) error {
		err := service.UpdateRoomByID(id, values)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteRoomByID() interface{} {
	return func(id uint) error {
		err := service.DeleteRoomByID(id)
		if err != nil {
			return err
		}
		return nil
	}
}
