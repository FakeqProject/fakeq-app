package model

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	GroupID     uint
	Name        string
	Avatar      string
	Description string
}
