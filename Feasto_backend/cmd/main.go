package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/riteshco/Feasto/pkg/api"
	"github.com/riteshco/Feasto/pkg/models"
)

func main(){

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading environment variables : %v" , err)
		return
	}

	_,err = models.InitDatabase()
	if err!= nil {
		log.Fatalf("Failed to initialize database : %v" , err)
	}

	api.Run()
}