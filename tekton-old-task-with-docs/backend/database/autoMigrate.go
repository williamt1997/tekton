package database

import (
	"fmt"
	"homers-backend/models"
	"log"

	"github.com/joho/godotenv"
)

func BuildDB() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error: Cannot Access ENV Files - ", err)
	}

	ConnectDB()
	err = DB.AutoMigrate(
		&models.PageView{},
	)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}

	SeedDatabase()
}

func SeedDatabase() {
	addPageView := models.PageView{
		ViewCount: 0,
	}
	DB.Create(&addPageView)
}
