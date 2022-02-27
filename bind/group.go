package bind

import (
	"github.com/FakeqProject/fakeq-app/dto"
	"github.com/FakeqProject/fakeq-app/model"
	"github.com/FakeqProject/fakeq-app/service"
	"github.com/webview/webview"
)

func GroupBindCollection(w webview.WebView) {
	w.Bind("getAllGroup", GetAllGroup())
	w.Bind("getAllGroupWithPagination", GetAllGroupWithPagination())
	w.Bind("getGroupByID", GetGroupByID())
	w.Bind("insertNewGroup", InsertNewGroup())
	w.Bind("updateGroupByID", UpdateGroupByID())
	w.Bind("deleteGroupByID", DeleteGroupByID())
}

func GetAllGroup() interface{} {
	return func() ([]dto.GroupResponseDto, error) {
		groups, err := service.GetAllGroup()
		if err != nil {
			return nil, err
		}
		var groupDtos []dto.GroupResponseDto
		for _, group := range groups {
			groupDtos = append(groupDtos, dto.CreateGroupResponseDto(group))
		}
		return groupDtos, nil
	}
}

func GetAllGroupWithPagination() interface{} {
	return func(page, limit int) ([]dto.GroupResponseDto, error) {
		groups, err := service.GetAllGroupWithPagination((page-1)*limit, limit)
		if err != nil {
			return nil, err
		}
		var groupDtos []dto.GroupResponseDto
		for _, group := range groups {
			groupDtos = append(groupDtos, dto.CreateGroupResponseDto(group))
		}
		return groupDtos, nil
	}
}

func GetGroupByID() interface{} {
	return func(id uint) (dto.GroupResponseDto, error) {
		group, err := service.GetGroupByID(id)
		if err != nil {
			return dto.GroupResponseDto{}, err
		}
		return dto.CreateGroupResponseDto(group), nil
	}
}

func InsertNewGroup() interface{} {
	return func(request dto.GroupRequestDto) (dto.GroupResponseDto, error) {
		group := model.Group{
			GroupID:     request.GroupID,
			Name:        request.Name,
			Avatar:      request.Avatar,
			Description: request.Description,
		}
		err := service.InsertNewGroup(group)
		if err != nil {
			return dto.GroupResponseDto{}, err
		}
		return dto.CreateGroupResponseDto(group), nil
	}
}

func UpdateGroupByID() interface{} {
	return func(id uint, values map[string]interface{}) error {
		err := service.UpdateGroupByID(id, values)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteGroupByID() interface{} {
	return func(id uint) error {
		err := service.DeleteGroupByID(id)
		if err != nil {
			return err
		}
		return nil
	}
}
