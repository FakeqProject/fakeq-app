package model

import "gorm.io/gorm"

type MessageFile struct {
	gorm.Model
	MessageID uint
	FileID    uint
}
