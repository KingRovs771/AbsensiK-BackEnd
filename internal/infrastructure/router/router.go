package router

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/delivery/http"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/delivery/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(userHandler *http.UserHandler,
	authHand *http.AuthHandler,
	departementHandler *http.DepartementHandler,
	roleHandler *http.RoleHandler,
	radiusHandler *http.RadiusHandler,
	schedulesHandler *http.SchedulesHandler,
	potonganHandler *http.PotonganHandler,
	tipePotonganHandler *http.TipePotonganHandler,
	izinHandler *http.IzinHandler,
	salaryhandler *http.SalaryHandler) *mux.Router {
	router := mux.NewRouter()

	//check API
	mainHandler := http.NewMainHandler()
	router.HandleFunc("/check", mainHandler.GetDomain).Methods("GET")

	// Get Profile Mobile

	//Login Auth
	router.HandleFunc("/v1/auth/login", authHand.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/auth/logout", authHand.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/auth/getInfo", authHand.GetUserInfo).Methods("GET", "OPTIONS")

	//Users
	router.HandleFunc("/v1/users/insertUsers", userHandler.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/users/allUsers", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/v1/users/updateUser/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/v1/users/deleteUser/{id}", userHandler.DeleteProfile).Methods("DELETE")
	router.HandleFunc("/v1/users/getUsersByIdUpdate/{id}", userHandler.GetUserByIdUpdate).Methods("GET")
	router.HandleFunc("/v1/users/getUsers/search", userHandler.SearchEmployeeByName).Methods("GET", "OPTIONS")

	//departements
	router.HandleFunc("/v1/departements/insertDepartements", departementHandler.CreateDepartements).Methods("POST")
	router.HandleFunc("/v1/departements/AllDepartements", departementHandler.GetAllDepartements).Methods("GET")
	router.HandleFunc("/v1/departements/getDepartementsById/{id}", departementHandler.GetDepartementsById).Methods("GET")
	router.HandleFunc("/v1/departements/updateDepartements/{id}", departementHandler.UpdateDepartement).Methods("PUT")
	router.HandleFunc("/v1/departements/deleteDepartements/{id}", departementHandler.DeleteDepartement).Methods("DELETE", "OPTIONS")

	//roles
	router.HandleFunc("/v1/roles/insertRoles", roleHandler.CreateRole).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/roles/AllRoles", roleHandler.GetAllRole).Methods("GET")
	router.HandleFunc("/v1/roles/getRolesById/{id}", roleHandler.GetRoleById).Methods("GET")
	router.HandleFunc("/v1/roles/updateRole/{id}", roleHandler.UpdateRole).Methods("PUT")
	router.HandleFunc("/v1/roles/deleteRole/{role_id}", roleHandler.DeleteRole).Methods("DELETE", "OPTIONS")

	//radius
	router.HandleFunc("/v1/radius/AllRadius", radiusHandler.GetAllRadius).Methods("GET")
	router.HandleFunc("/v1/radius/insertRadius", radiusHandler.CreateRadius).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/radius/getRadiusById/{id:[1-9]+}", radiusHandler.GetRadiusById).Methods("GET")
	router.HandleFunc("/v1/radius/updateRaidus/{id:[1-9]+}", radiusHandler.UpdateRadius).Methods("PUT")
	router.HandleFunc("/v1/radius/deleteRadius/{id:[1-9]+}", radiusHandler.DeleteRadius).Methods("DELETE")

	//schedules
	router.HandleFunc("/v1/schdules/allSchedules", schedulesHandler.GetAllSchedules).Methods("GET")
	router.HandleFunc("/v1/schdules/insertSchedules", schedulesHandler.CreateSchedules).Methods("POST")
	router.HandleFunc("/v1/schdules/getSchedulesById/{id:[1-9]+}", schedulesHandler.GetSchedulesById).Methods("GET")
	router.HandleFunc("/v1/schdules/updateSchedules/{id:[1-9]+}", schedulesHandler.UpdateSchedules).Methods("PUT")
	router.HandleFunc("/v1/schdules/deleteSchedules/{schedule_id}", schedulesHandler.DeleteSchedules).Methods("DELETE", "OPTIONS")

	//tipe Potongan
	router.HandleFunc("/v1/tipePotongan/allTipePotongan", tipePotonganHandler.GetAllTipePotongan).Methods("GET")
	router.HandleFunc("/v1/tipePotongan/insertTipePotongan", tipePotonganHandler.CreateTipePotongan).Methods("POST")
	router.HandleFunc("/v1/tipePotongan/getTipeById/{id:[1-9]+}", tipePotonganHandler.GetTipePotonganById).Methods("GET")
	router.HandleFunc("/v1/tipePotongan/updateTipePotongan/{id:[1-9]+}", tipePotonganHandler.UpdateTipePotongan).Methods("PUT")
	router.HandleFunc("/v1/tipePotongan/deleteTipePotongan/{id:[1-9]+}", tipePotonganHandler.DeleteTipePotongan).Methods("DELETE")

	//izin
	router.HandleFunc("/v1/izin/allIzin", izinHandler.GetAllIzin).Methods("GET")
	router.HandleFunc("/v1/izin/insertIzin", izinHandler.CreateIzin).Methods("POST")
	router.HandleFunc("/v1/izin/getIzinById", izinHandler.GetIzinById).Methods("GET")
	router.HandleFunc("/v1/izin/updateIzin", izinHandler.UpdateIzin).Methods("PUT")
	router.HandleFunc("/v1/izin/deleteIzin", izinHandler.DeleteIzin).Methods("DELETE")
	router.HandleFunc("/v1/izin/{id}/approve", izinHandler.ApproveIzin).Methods("PUT", "OPTIONS")

	//salary
	router.HandleFunc("/v1/salary/allSalary", salaryhandler.GetSalariesByMonth).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/salary/allSalaryAndName", salaryhandler.GetSalariesByMonthAndName).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/salary/checkAPI", salaryhandler.GetSalary).Methods("GET")
	router.HandleFunc("/v1/salary/insertSalary", salaryhandler.GenerateSalary).Methods("POST", "OPTIONS")

	//return
	router.Use(middleware.CORSMiddleware())
	return router
}
