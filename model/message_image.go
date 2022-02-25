package model

import "gorm.io/gorm"

type MessageImage struct {
	gorm.Model
	MessageID uint
	ImageID   uint
}
