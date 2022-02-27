package bind

import (
	"github.com/FakeqProject/fakeq-app/dto"
	"github.com/FakeqProject/fakeq-app/model"
	"github.com/FakeqProject/fakeq-app/service"
	"github.com/webview/webview"
)

func MessageBindCollection(w webview.WebView) {
	w.Bind("getAllMessage", GetAllMessage())
	w.Bind("getAllMessageWithPagination", GetAllMessageWithPagination())
	w.Bind("getMessageByID", GetMessageByID())
	w.Bind("insertNewMessage", InsertNewMessage())
	w.Bind("updateMessageByID", UpdateMessageByID())
	w.Bind("deleteMessageByID", DeleteMessageByID())
}

func GetAllMessage() interface{} {
	return func() ([]dto.MessageResponseDto, error) {
		messages, err := service.GetAllMessage()
		if err != nil {
			return nil, err
		}
		var messageDtos []dto.MessageResponseDto
		for _, message := range messages {
			messageDtos = append(messageDtos, dto.CreateMessageResponseDto(message))
		}
		return messageDtos, nil
	}
}

func GetAllMessageWithPagination() interface{} {
	return func(page, limit int) ([]dto.MessageResponseDto, error) {
		messages, err := service.GetAllMessageWithPagination((page-1)*limit, limit)
		if err != nil {
			return nil, err
		}
		var messageDtos []dto.MessageResponseDto
		for _, message := range messages {
			messageDtos = append(messageDtos, dto.CreateMessageResponseDto(message))
		}
		return messageDtos, nil
	}
}

func GetMessageByID() interface{} {
	return func(id uint) (dto.MessageResponseDto, error) {
		message, err := service.GetMessageByID(id)
		if err != nil {
			return dto.MessageResponseDto{}, err
		}
		return dto.CreateMessageResponseDto(message), nil
	}
}

func InsertNewMessage() interface{} {
	return func(request dto.MessageRequestDto) (dto.MessageResponseDto, error) {
		// TODO: insert files and images
		message := model.Message{
			SenderID:     request.SenderID,
			SenderName:   request.SenderName,
			Content:      request.Content,
			Code:         request.Code,
			Date:         request.Date,
			UnixTime:     request.UnixTime,
			ReplyMessage: request.ReplyMessage,
			Recalled:     request.Recalled,
			RoomId:       request.RoomId,
			// Files:        request.Files,
			// Images:       request.Images,
		}
		err := service.InsertNewMessage(message)
		if err != nil {
			return dto.MessageResponseDto{}, err
		}
		return dto.CreateMessageResponseDto(message), nil
	}
}

func UpdateMessageByID() interface{} {
	return func(id uint, values map[string]interface{}) error {
		err := service.UpdateMessageByID(id, values)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteMessageByID() interface{} {
	return func(id uint) error {
		err := service.DeleteMessageByID(id)
		if err != nil {
			return err
		}
		return nil
	}
}
