package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageCDN  string
	ImageURL  string
	ImageSize int
}
