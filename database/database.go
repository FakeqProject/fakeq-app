package database

import "gorm.io/gorm"

type DatabaseInterface interface {
	InitDatabase(interface{}) error
	GetDB() *gorm.DB
}
