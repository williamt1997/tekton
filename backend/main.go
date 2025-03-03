package main

import (
	"fmt"
	"homers-backend/database"
	"homers-backend/router"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error: Cannot Access ENV Files - ", err)
	}
	database.BuildDB()
	database.ConnectDB()
}

func main() {
	router.Start()
}
