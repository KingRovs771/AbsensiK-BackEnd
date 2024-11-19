package main

import (
	"log"
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/app"
	"github.com/joho/godotenv"
)


func main(){

	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error Loading .env file")
	}
	
	apps := app.NewApp()
	router := apps.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}