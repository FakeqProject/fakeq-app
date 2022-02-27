package dto

import "github.com/FakeqProject/fakeq-app/model"

type UserResponseDto struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	Slogan   string `json:"slogan"`
}

type UserRequestDto struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	Avatar   string `json:"avatar"`
	Slogan   string `json:"slogan"`
}

func CreateUserResponseDto(user model.User) UserResponseDto {
	return UserResponseDto{
		ID:       user.ID,
		UserID:   user.UserID,
		UserName: user.UserName,
		Avatar:   user.Avatar,
		Slogan:   user.Slogan,
	}
}
