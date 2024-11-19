package router

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/delivery/http"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *http.UserHandler) *mux.Router{
	router := mux.NewRouter()

	mainHandler := http.NewMainHandler()
	router.HandleFunc("/check", mainHandler.GetDomain).Methods("GET")

	http.UserHandlres(router, userHandler)
	return router
}