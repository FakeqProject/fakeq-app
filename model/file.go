package model

import "gorm.io/gorm"

type File struct {
	gorm.Model
	FileCDN  string
	FileURL  string
	FileSize int
}
