package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {

	url := "root:root@tcp(localhost:3306)/UserAuthDB?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}
