package models

import "time"

type User struct {
	ID        uint64 `gorm:"primarykey"`
	Username  string `gorm:"size:64"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}