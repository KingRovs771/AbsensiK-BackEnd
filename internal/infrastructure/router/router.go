package router

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/delivery/http"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *http.UserHandler, authHand *http.AuthHandler, departementHandler *http.DepartementHandler) *mux.Router {
	router := mux.NewRouter()

	//check API
	mainHandler := http.NewMainHandler()
	router.HandleFunc("/check", mainHandler.GetDomain).Methods("GET")

	// Get Profile Mobile
	router.HandleFunc("/users/getProfile", userHandler.GetProfile).Methods("GET")

	//Login  Auth
	router.HandleFunc("/v1/auth/login", authHand.Login).Methods("POST")

	//Users
	router.HandleFunc("/v1/users/insertUsers", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/v1/users/allUsers", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/v1/users/updateUser/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/v1/users/deleteUser/{id}", userHandler.DeleteProfile).Methods("DELETE")
	router.HandleFunc("/v1/users/getUsersByIdUpdate/{id}", userHandler.GetUserByIdUpdate).Methods("GET")

	//departements
	router.HandleFunc("/v1/departements/insertDepartements", departementHandler.CreateDepartements).Methods("POST")
	router.HandleFunc("/v1/departements/getAllDepartements", departementHandler.GetAllDepartements).Methods("GET")
	router.HandleFunc("/v1/departements/getDepartementsById/{id}", departementHandler.GetDepartementsById).Methods("GET")
	router.HandleFunc("/v1/departements/updateDepartements/{id}", departementHandler.UpdateDepartement).Methods("PUT")
	router.HandleFunc("/v1/departements/deleteDepartements/{id}", departementHandler.DeleteDepartement).Methods("DELETE")

	return router
}
