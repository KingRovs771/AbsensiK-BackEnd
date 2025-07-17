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
	salaryhandler *http.SalaryHandler,
	faceHandler *http.FaceHandler,
	attendanceHandler *http.AttendanceHandler,
	dashboardHandler *http.DashboardHandler,
	reportHandler *http.ReportHandler) *mux.Router {
	router := mux.NewRouter()

	//check API
	mainHandler := http.NewMainHandler()
	router.HandleFunc("/check", mainHandler.GetDomain).Methods("GET")

	//Dashbaord
	router.HandleFunc("/v1/dashboard/status", dashboardHandler.GetDashboardStats).Methods("GET")
	// Get Profile Mobile

	//Login Auth
	router.HandleFunc("/v1/auth/login", authHand.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/auth/logout", authHand.Logout).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/auth/getInfo", authHand.GetUserInfo).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/auth/profile", authHand.ProfileMobile).Methods("GET")

	//Users
	router.HandleFunc("/v1/users/insertUsers", userHandler.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/users/allUsers", userHandler.GetAllUsers).Methods("GET")
	router.HandleFunc("/v1/users/updateUser/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/v1/users/deleteUser/{id}", userHandler.DeleteProfile).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/v1/users/getUsersByIdUpdate/{id}", userHandler.GetUserByIdUpdate).Methods("GET")
	router.HandleFunc("/v1/users/getUsers/search", userHandler.SearchEmployeeByName).Methods("GET", "OPTIONS")

	//departements
	router.HandleFunc("/v1/departements/insertDepartements", departementHandler.CreateDepartements).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/departements/AllDepartements", departementHandler.GetAllDepartements).Methods("GET")
	router.HandleFunc("/v1/departements/getDepartementsById/{id}", departementHandler.GetDepartementsById).Methods("GET")
	router.HandleFunc("/v1/departements/updateDepartements/{id}", departementHandler.UpdateDepartement).Methods("PUT", "OPTIONS")
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
	router.HandleFunc("/v1/radius/deleteRadius/{id:[1-9]+}", radiusHandler.DeleteRadius).Methods("DELETE", "OPTIONS")

	//schedules
	router.HandleFunc("/v1/schdules/allSchedules", schedulesHandler.GetAllSchedules).Methods("GET")
	router.HandleFunc("/v1/schdules/insertSchedules", schedulesHandler.CreateSchedules).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/schdules/getSchedulesById/{id:[1-9]+}", schedulesHandler.GetSchedulesById).Methods("GET")
	router.HandleFunc("/v1/schdules/updateSchedules/{id:[1-9]+}", schedulesHandler.UpdateSchedules).Methods("PUT")
	router.HandleFunc("/v1/schdules/deleteSchedules/{schedule_id}", schedulesHandler.DeleteSchedules).Methods("DELETE", "OPTIONS")

	//tipe Potongan
	router.HandleFunc("/v1/tipePotongan/allTipePotongan", tipePotonganHandler.GetAllTipePotongan).Methods("GET")
	router.HandleFunc("/v1/tipePotongan/insertTipePotongan", tipePotonganHandler.CreateTipePotongan).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/tipePotongan/getTipeById/{tipe_potongan_id}", tipePotonganHandler.GetTipePotonganById).Methods("GET")
	router.HandleFunc("/v1/tipePotongan/updateTipePotongan/{tipe_potongan_id}", tipePotonganHandler.UpdateTipePotongan).Methods("PUT")
	router.HandleFunc("/v1/tipePotongan/deleteTipePotongan/{tipe_potongan_id}", tipePotonganHandler.DeleteTipePotongan).Methods("DELETE", "OPTIONS")

	//Potongan
	router.HandleFunc("/v1/potongan/allPotongan", potonganHandler.GetAllPotongan).Methods("GET")
	router.HandleFunc("/v1/potongan/insertPotongan", potonganHandler.CreatePotongan).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/potongan/getPotonganById/{potongan_id}", potonganHandler.GetPotonganById).Methods("GET")
	router.HandleFunc("/v1/potongan/updatePotongan/{potongan_id}", potonganHandler.UpdatePotongan).Methods("PUT")
	router.HandleFunc("/v1/potongan/deletePotongan/{potongan_id}", potonganHandler.UpdatePotongan).Methods("DELETE", "OPTIONS")

	//izin
	router.HandleFunc("/v1/izin/allIzin", izinHandler.GetAllIzin).Methods("GET")
	router.HandleFunc("/v1/izin/insertIzin", izinHandler.CreateIzin).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/izin/getIzinById", izinHandler.GetIzinById).Methods("GET")
	router.HandleFunc("/v1/izin/updateIzin", izinHandler.UpdateIzin).Methods("PUT")
	router.HandleFunc("/v1/izin/deleteIzin", izinHandler.DeleteIzin).Methods("DELETE")
	router.HandleFunc("/v1/izin/{id}/approve", izinHandler.ApproveIzin).Methods("PUT", "OPTIONS")
	router.HandleFunc("/v1/izin/getIzinByuserUID", izinHandler.GetUserPermitHistory).Methods("GET")

	//salary
	router.HandleFunc("/v1/salary/allSalary", salaryhandler.GetSalariesByMonth).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/salary/allSalaryAndName", salaryhandler.GetSalariesByMonthAndName).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/salary/checkAPI", salaryhandler.GetSalary).Methods("GET")
	router.HandleFunc("/v1/salary/insertSalary", salaryhandler.GenerateSalary).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/salary/getLatestSalary", salaryhandler.GetLatestPayslip).Methods("GET", "OPTIONS")

	//Face Upload
	router.HandleFunc("/v1/face/uploadFoto", faceHandler.UploadFaceHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/face/getAllFoto", faceHandler.GetAllFacesHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/face/deleteFoto/{id}", faceHandler.DeleteFoto).Methods("DELETE", "OPTIONS")

	// Attendances
	router.HandleFunc("/v1/attendances/getDataAttendances", attendanceHandler.GetAttendanceData).Methods("GET", "OPTIONS")
	router.HandleFunc("/v1/attendances/clockIn", attendanceHandler.ClockIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/attendances/clockOut", attendanceHandler.ClockOut).Methods("POST", "OPTIONS")
	router.HandleFunc("/v1/attendances/getAllAttendances", attendanceHandler.GetAllAttendances).Methods("GET", "OPTIONS")

	//Report
	router.HandleFunc("/v1/report/generateReport", reportHandler.GenerateReport).Methods("GET", "OPTIONS")
	//return
	router.Use(middleware.CORSMiddleware())
	return router
}
