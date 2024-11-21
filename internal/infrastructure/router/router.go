package router

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/delivery/http"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *http.UserHandler, authHand *http.AuthHandler) *mux.Router{
	router := mux.NewRouter()

	mainHandler := http.NewMainHandler()
	router.HandleFunc("/check", mainHandler.GetDomain).Methods("GET")

	http.UserHandlres(router, userHandler)

	//Login  Auth
	router.HandleFunc("/v1/auth/login", authHand.Login).Methods("POST")
	return router
}