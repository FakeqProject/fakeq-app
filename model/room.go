package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	RoomID uint
	Name   string
	Type   string
}
