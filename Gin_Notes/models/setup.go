package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("notes:tmp_pwd@tcp(localhost:3306)/notes?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database : " + err.Error())
	}
	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Note{})
}
