package bind

import (
	"github.com/FakeqProject/fakeq-app/dto"
	"github.com/FakeqProject/fakeq-app/model"
	"github.com/FakeqProject/fakeq-app/service"
	"github.com/webview/webview"
)

func UserBindCollection(w webview.WebView) {
	w.Bind("getAllUser", GetAllUser())
	w.Bind("getAllUserWithPagination", GetAllUserWithPagination())
	w.Bind("getUserByID", GetUserByID())
	w.Bind("insertNewUser", InsertNewUser())
	w.Bind("updateUserByID", UpdateUserByID())
	w.Bind("deleteUserByID", DeleteUserByID())
}

func GetAllUser() interface{} {
	return func() ([]dto.UserResponseDto, error) {
		users, err := service.GetAllUser()
		if err != nil {
			return nil, err
		}
		var userDtos []dto.UserResponseDto
		for _, user := range users {
			userDtos = append(userDtos, dto.CreateUserResponseDto(user))
		}
		return userDtos, nil
	}
}

func GetAllUserWithPagination() interface{} {
	return func(page, limit int) ([]dto.UserResponseDto, error) {
		users, err := service.GetAllUserWithPagination((page-1)*limit, limit)
		if err != nil {
			return nil, err
		}
		var userDtos []dto.UserResponseDto
		for _, user := range users {
			userDtos = append(userDtos, dto.CreateUserResponseDto(user))
		}
		return userDtos, nil
	}
}

func GetUserByID() interface{} {
	return func(id uint) (dto.UserResponseDto, error) {
		user, err := service.GetUserByID(id)
		if err != nil {
			return dto.UserResponseDto{}, err
		}
		return dto.CreateUserResponseDto(user), nil
	}
}

func InsertNewUser() interface{} {
	return func(request dto.UserRequestDto) (dto.UserResponseDto, error) {
		user := model.User{
			UserID:   request.UserID,
			UserName: request.UserName,
			Avatar:   request.Avatar,
			Slogan:   request.Slogan,
		}
		err := service.InsertNewUser(user)
		if err != nil {
			return dto.UserResponseDto{}, err
		}
		return dto.CreateUserResponseDto(user), nil
	}
}

func UpdateUserByID() interface{} {
	return func(id uint, values map[string]interface{}) error {
		err := service.UpdateUserByID(id, values)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteUserByID() interface{} {
	return func(id uint) error {
		err := service.DeleteUserByID(id)
		if err != nil {
			return err
		}
		return nil
	}
}
