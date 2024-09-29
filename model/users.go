package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
}

func SeedUsers(db *gorm.DB) {
	users := []*User{
		{Name: "Rocky Racoon"},
		{Name: "Another User"},
		{Name: "Test User"},
	}
	db.Create(users)
}
