package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SenderID     uint
	SenderName   string
	Content      string
	Code         string
	Date         string
	UnixTime     int64
	ReplyMessage string
	Recalled     bool
	RoomID       uint
}
