package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   uint
	UserName string
	Avatar   string
	Slogan   string
}
