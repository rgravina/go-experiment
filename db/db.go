package db

import (
	"github.com/glebarez/sqlite"
	"go-experiment/model"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&model.User{})
	model.SeedUsers(db)
}

func DbManager() *gorm.DB {
	return db
}
