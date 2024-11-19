package http

import "github.com/gorilla/mux"

func UserHandlres(r *mux.Router, userHandler *UserHandler){
	r.HandleFunc("/users/check", userHandler.getCheckUsers).Methods("GET")
	r.HandleFunc("/users/insertUsers", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/AllUsers", userHandler.GetAllUsers).Methods("GET")
}

func AuthHandlers(r *mux.Router, authHandlre *AuthHandler){
	
}