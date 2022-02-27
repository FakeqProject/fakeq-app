package dto

import "github.com/FakeqProject/fakeq-app/model"

type GroupResponseDto struct {
	ID          uint   `json:"id"`
	GroupID     uint   `json:"group_id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type GroupRequestDto struct {
	GroupID     uint   `json:"group_id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

func CreateGroupResponseDto(group model.Group) GroupResponseDto {
	return GroupResponseDto{
		ID:          group.ID,
		GroupID:     group.GroupID,
		Name:        group.Name,
		Avatar:      group.Avatar,
		Description: group.Description,
	}
}
