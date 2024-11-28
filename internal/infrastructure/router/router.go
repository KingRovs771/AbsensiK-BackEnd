package router

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/delivery/http"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *http.UserHandler,
	authHand *http.AuthHandler,
	departementHandler *http.DepartementHandler,
	roleHandler *http.RoleHandler,
	radiusHandler *http.RadiusHandler,
	schedulesHandler *http.SchedulesHandler) *mux.Router {
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
	router.HandleFunc("/v1/departements/AllDepartements", departementHandler.GetAllDepartements).Methods("GET")
	router.HandleFunc("/v1/departements/getDepartementsById/{id}", departementHandler.GetDepartementsById).Methods("GET")
	router.HandleFunc("/v1/departements/updateDepartements/{id}", departementHandler.UpdateDepartement).Methods("PUT")
	router.HandleFunc("/v1/departements/deleteDepartements/{id}", departementHandler.DeleteDepartement).Methods("DELETE")

	//roles
	router.HandleFunc("/v1/roles/insertRoles", roleHandler.CreateRole).Methods("POST")
	router.HandleFunc("/v1/roles/AllRoles", roleHandler.GetAllRole).Methods("GET")
	router.HandleFunc("/v1/roles/getRolesById/{id}", roleHandler.GetRoleById).Methods("GET")
	router.HandleFunc("/v1/roles/updateRole/{id}", roleHandler.UpdateRole).Methods("PUT")
	router.HandleFunc("/v1/roles/deleteRole/{id}", roleHandler.DeleteRole).Methods("DELETE")

	//radius
	router.HandleFunc("/v1/radius/AllRadius", radiusHandler.GetAllRadius).Methods("GET")
	router.HandleFunc("/v1/radius/insertRadius", radiusHandler.CreateRadius).Methods("POST")
	router.HandleFunc("/v1/radius/getRadiusById/{id:[1-9]+}", radiusHandler.GetRadiusById).Methods("GET")
	router.HandleFunc("/v1/radius/updateRaidus/{id:[1-9]+}", radiusHandler.UpdateRadius).Methods("PUT")
	router.HandleFunc("/v1/radius/deleteRadius/{id:[1-9]+}", radiusHandler.DeleteRadius).Methods("DELETE")

	//schedules
	router.HandleFunc("v1/schdules/allSchedules", schedulesHandler.GetAllSchedules).Methods("GET")
	router.HandleFunc("v1/schdules/insertSchedules", schedulesHandler.CreateSchedules).Methods("GET")
	router.HandleFunc("v1/schdules/getSchedulesById/{id:[1-9]+}", schedulesHandler.GetSchedulesById).Methods("GET")
	router.HandleFunc("v1/schdules/updateSchedules/{id:[1-9]+}", schedulesHandler.UpdateSchedules).Methods("PUT")
	router.HandleFunc("v1/schdules/deleteSchedules/{id:[1-9]+}", schedulesHandler.DeleteSchedules).Methods("DELETE")

	//return
	return router
}
