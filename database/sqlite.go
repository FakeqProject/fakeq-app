package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDatabase struct {
	FilePath string
}

func (s SqliteDatabase) InitDB(migrateDst ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s.FilePath), &gorm.Config{})
	if err != nil {
		return db, err
	}
	db.AutoMigrate(migrateDst...)

	return db, nil
}
