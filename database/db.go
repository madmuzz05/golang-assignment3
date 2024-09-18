package database

import (
	"fmt"
	"log"
	"update-data/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	USER     = "localhost"
	PASSWORD = "password"
	PORT     = "5432"
	DATABASE = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE, PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Success connected to database")
	db.Debug().AutoMigrate(models.Log{})
}

func GetDB() *gorm.DB {
	return db
}
