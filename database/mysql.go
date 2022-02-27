package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDatabase struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	CharSet  string `json:"charset"`
}

func (m MysqlDatabase) InitDB(migrateDst ...interface{}) (*gorm.DB, error) {
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		m.UserName,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
		m.CharSet)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		return db, err
	}
	db.AutoMigrate(migrateDst...)

	return db, nil
}
