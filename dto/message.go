package dto

import (
	"github.com/FakeqProject/fakeq-app/model"
	"github.com/FakeqProject/fakeq-app/service"
)

type MessageResponseDto struct {
	ID           uint   `json:"id"`
	SenderID     uint   `json:"sender_id"`
	SenderName   string `json:"sender_name"`
	Content      string `json:"content"`
	Code         string `json:"code"`
	Date         string `json:"date"`
	UnixTime     int64  `json:"unix_time"`
	ReplyMessage string `json:"reply_message"`
	Recalled     bool   `json:"recalled"`
	RoomID       uint   `json:"room_id"`
	FileIDs      []uint `json:"file_ids"`
	ImageIDs     []uint `json:"image_ids"`
}

type MessageRequestDto struct {
	SenderID     uint   `json:"sender_id"`
	SenderName   string `json:"sender_name"`
	Content      string `json:"content"`
	Code         string `json:"code"`
	Date         string `json:"date"`
	UnixTime     int64  `json:"unix_time"`
	ReplyMessage string `json:"reply_message"`
	Recalled     bool   `json:"recalled"`
	RoomID       uint   `json:"room_id"`
}

func CreateMessageResponseDto(message model.Message) (MessageResponseDto, error) {
	messageFiles, err := service.GetAllMessageFileByMessageID(message.ID)
	if err != nil {
		return MessageResponseDto{}, err
	}
	messageImages, err := service.GetAllMessageImageByMessageID(message.ID)
	if err != nil {
		return MessageResponseDto{}, err
	}
	fileIDs := make([]uint, len(messageFiles))
	imageIDs := make([]uint, len(messageImages))
	for i, messageFile := range messageFiles {
		fileIDs[i] = messageFile.ID
	}
	for i, messageImage := range messageImages {
		imageIDs[i] = messageImage.ID
	}
	return MessageResponseDto{
		ID:           message.ID,
		SenderID:     message.SenderID,
		SenderName:   message.SenderName,
		Content:      message.Content,
		Code:         message.Code,
		Date:         message.Date,
		UnixTime:     message.UnixTime,
		ReplyMessage: message.ReplyMessage,
		Recalled:     message.Recalled,
		RoomID:       message.RoomID,
		FileIDs:      fileIDs,
		ImageIDs:     imageIDs,
	}, nil
}
