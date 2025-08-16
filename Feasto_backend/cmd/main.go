package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/riteshco/Feasto/pkg/api"
	"github.com/riteshco/Feasto/pkg/models"
)

func main(){

	if err := godotenv.Load(); err != nil {
		log.Println("INFO: .env file not found, using system environment variables.")
	}

	_,err := models.InitDatabase()
	if err!= nil {
		log.Fatalf("Failed to initialize database : %v" , err)
	}

	api.Run()
}