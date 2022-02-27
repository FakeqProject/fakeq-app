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
	w.Bind("getAllMessageFileByMessageID", GetAllMessageFileByMessageID())
	w.Bind("getAllMessageImageByMessageID", GetAllMessageImageByMessageID())
	w.Bind("insertNewMessageFileByMessageID", InsertNewMessageFileByMessageID())
	w.Bind("insertNewMessageImageByMessageID", InsertNewMessageImageByMessageID())
	w.Bind("deleteMessageFileByMessageIDAndFileID", DeleteMessageFileByMessageIDAndFileID())
	w.Bind("deleteMessageImageByMessageIDAndImageID", DeleteMessageImageByMessageIDAndImageID())
}

func GetAllMessage() interface{} {
	return func() ([]dto.MessageResponseDto, error) {
		messages, err := service.GetAllMessage()
		if err != nil {
			return nil, err
		}
		var messageDtos []dto.MessageResponseDto
		for _, message := range messages {
			messageDto, err := dto.CreateMessageResponseDto(message)
			if err != nil {
				return nil, err
			}
			messageDtos = append(messageDtos, messageDto)
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
			messageDto, err := dto.CreateMessageResponseDto(message)
			if err != nil {
				return nil, err
			}
			messageDtos = append(messageDtos, messageDto)
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
		messageDto, err := dto.CreateMessageResponseDto(message)
		if err != nil {
			return dto.MessageResponseDto{}, err
		}

		return messageDto, nil
	}
}

func InsertNewMessage() interface{} {
	return func(request dto.MessageRequestDto) (dto.MessageResponseDto, error) {
		message := model.Message{
			SenderID:     request.SenderID,
			SenderName:   request.SenderName,
			Content:      request.Content,
			Code:         request.Code,
			Date:         request.Date,
			UnixTime:     request.UnixTime,
			ReplyMessage: request.ReplyMessage,
			Recalled:     request.Recalled,
			RoomID:       request.RoomID,
		}
		err := service.InsertNewMessage(message)
		if err != nil {
			return dto.MessageResponseDto{}, err
		}
		messageDto, err := dto.CreateMessageResponseDto(message)
		if err != nil {
			return dto.MessageResponseDto{}, err
		}
		return messageDto, nil
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

func GetAllMessageFileByMessageID() interface{} {
	return func(id uint) ([]dto.FileResponseDto, error) {
		messageFiles, err := service.GetAllMessageFileByMessageID(id)
		if err != nil {
			return nil, err
		}
		var fileDtos []dto.FileResponseDto
		for _, messageFile := range messageFiles {
			file, err := service.GetFileByID(messageFile.FileID)
			if err != nil {
				return nil, err
			}
			fileDto := dto.CreateFileResponseDto(file)
			fileDtos = append(fileDtos, fileDto)
		}
		return fileDtos, nil
	}
}

func GetAllMessageImageByMessageID() interface{} {
	return func(id uint) ([]dto.ImageResponseDto, error) {
		messageImages, err := service.GetAllMessageImageByMessageID(id)
		if err != nil {
			return nil, err
		}
		var imageDtos []dto.ImageResponseDto
		for _, messageImage := range messageImages {
			image, err := service.GetImageByID(messageImage.ImageID)
			if err != nil {
				return nil, err
			}
			imageDto := dto.CreateImageResponseDto(image)
			imageDtos = append(imageDtos, imageDto)
		}
		return imageDtos, nil
	}
}

func InsertNewMessageFileByMessageID() interface{} {
	return func(messageID, fileID uint) error {
		messageFile := model.MessageFile{
			MessageID: messageID,
			FileID:    fileID,
		}
		err := service.InsertNewMessageFile(messageFile)
		if err != nil {
			return err
		}
		return nil
	}
}

func InsertNewMessageImageByMessageID() interface{} {
	return func(messageID, fileID uint) error {
		messageImage := model.MessageImage{
			MessageID: messageID,
			ImageID:   fileID,
		}
		err := service.InsertNewMessageImage(messageImage)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteMessageFileByMessageIDAndFileID() interface{} {
	return func(messageID, fileID uint) error {
		err := service.DeleteMessageFileByMessageIDAndFileID(messageID, fileID)
		if err != nil {
			return err
		}
		return nil
	}
}

func DeleteMessageImageByMessageIDAndImageID() interface{} {
	return func(messageID, imageID uint) error {
		err := service.DeleteMessageImageByMessageIDAndImageID(messageID, imageID)
		if err != nil {
			return err
		}
		return nil
	}
}
