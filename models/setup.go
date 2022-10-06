package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/db_restapi_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
