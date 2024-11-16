package main

import (
	"log"
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/app"
)


func main(){
	apps := app.NewApp()
	router := apps.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}