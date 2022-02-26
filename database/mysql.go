package database

import (
	"fmt"

	"github.com/FakeqProject/fakeq-app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

type MysqlDatabase struct {
}

type MysqlConfig struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	CharSet  string `json:"charset"`
}

func (m MysqlDatabase) InitDatabase(config interface{}) error {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		config.(MysqlConfig).UserName,
		config.(MysqlConfig).Password,
		config.(MysqlConfig).Host,
		config.(MysqlConfig).Port,
		config.(MysqlConfig).Database,
		config.(MysqlConfig).CharSet)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
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

	MysqlDB = db
	return nil
}

func (m MysqlDatabase) GetDB() *gorm.DB {
	return MysqlDB
}
