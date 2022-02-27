package database

import (
	"github.com/FakeqProject/fakeq-app/model"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DatabaseInterface interface {
	InitDB(migrateDst ...interface{}) (*gorm.DB, error)
}

func InitDatabase(dbi DatabaseInterface) {
	db, err := dbi.InitDB(
		&model.File{},
		&model.Group{},
		&model.Image{},
		&model.MessageFile{},
		&model.MessageImage{},
		&model.Message{},
		&model.Room{},
		&model.User{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
