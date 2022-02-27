package dto

import "github.com/FakeqProject/fakeq-app/model"

type MessageResponseDto struct {
	ID           uint               `json:"id"`
	SenderID     uint               `json:"sender_id"`
	SenderName   string             `json:"sender_name"`
	Content      string             `json:"content"`
	Code         string             `json:"code"`
	Date         string             `json:"date"`
	UnixTime     int64              `json:"unix_time"`
	ReplyMessage string             `json:"reply_message"`
	Recalled     bool               `json:"recalled"`
	RoomId       uint               `json:"room_id"`
	Files        []FileResponseDto  `json:"files"`
	Images       []ImageResponseDto `json:"images"`
}

type MessageRequestDto struct {
	SenderID     uint              `json:"sender_id"`
	SenderName   string            `json:"sender_name"`
	Content      string            `json:"content"`
	Code         string            `json:"code"`
	Date         string            `json:"date"`
	UnixTime     int64             `json:"unix_time"`
	ReplyMessage string            `json:"reply_message"`
	Recalled     bool              `json:"recalled"`
	RoomId       uint              `json:"room_id"`
	Files        []FileRequestDto  `json:"files"`
	Images       []ImageRequestDto `json:"images"`
}

func CreateMessageResponseDto(message model.Message) MessageResponseDto {
	// TODO : add files and images
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
		RoomId:       message.RoomId,
		// Files:        CreateFileResponseDtoSlice(message.Files),
		// Images:       CreateImageResponseDtoSlice(message.Images),
	}
}
