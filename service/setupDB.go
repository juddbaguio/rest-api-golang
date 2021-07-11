package service

import (
	"fmt"
	"log"
	"os"

	"github.com/juddbaguio/rest-api-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeDB() *gorm.DB {
	// load Env Variables
	HOST := os.Getenv("HOST")
	DB_PORT := os.Getenv("DB_PORT")
	USER := os.Getenv("USER")
	NAME := os.Getenv("NAME")
	PASSWORD := os.Getenv("PASSWORD")

	// Data connection string
	DB_URI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", HOST, USER, NAME, PASSWORD, DB_PORT)
	
	// Open DB
	db, err := gorm.Open(postgres.Open(DB_URI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connected successfully")
	}

	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Book{})

	return db
}