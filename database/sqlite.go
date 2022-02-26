package database

import (
	"github.com/FakeqProject/fakeq-app/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SqliteDB *gorm.DB

type SqliteDatabase struct {
}

type SqliteConfig struct {
	FilePath string `json:"filepath"`
}

func (s SqliteDatabase) InitDatabase(config interface{}) error {
	db, err := gorm.Open(sqlite.Open(config.(SqliteConfig).FilePath), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&model.File{})
	db.AutoMigrate(&model.Group{})
	db.AutoMigrate(&model.Image{})
	db.AutoMigrate(&model.MessageFile{})
	db.AutoMigrate(&model.MessageImage{})
	db.AutoMigrate(&model.Message{})
	db.AutoMigrate(&model.Room{})
	db.AutoMigrate(&model.User{})

	SqliteDB = db
	return nil
}

func (s SqliteDatabase) GetDB() *gorm.DB {
	return SqliteDB
}
