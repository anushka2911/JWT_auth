package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	url := "root:root@tcp(localhost:3306)/UserAuthDB?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	DB = db
}
