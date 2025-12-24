package models

import (
	"notes_application/helpers"
	"time"
)

type User struct {
	ID        uint64 `gorm:"primarykey"`
	Username  string `gorm:"size:64"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserCheckAvailability(email string) bool {
	var user User
	DB.Where("username = ?", email).First(&user)
	return (user.ID == 0)
}

func UserCreate(email string, password string) *User {
	hshPasswd, _ := helpers.HashPassword(password)
	user := User{
		Username: email,
		Password: hshPasswd,
	}
	DB.Create(&user)
	return &user
}

func UserFind(id uint64) *User {
	var user User
	DB.Where("id = ?", id).First(&user)
	return &user
}

func UserCheck(email string, password string) *User {
	var user User
	DB.Where("username = ?", email).First(&user)
	if user.ID == 0 {
		return nil
	}
	if helpers.CheckPasswordHash(password, user.Password) {
		return &user
	}
	return nil
}
