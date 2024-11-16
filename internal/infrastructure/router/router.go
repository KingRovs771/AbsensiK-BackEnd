package router

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/delivery/http"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *http.UserHandler) *mux.Router{
	router := mux.NewRouter()
	return router
}