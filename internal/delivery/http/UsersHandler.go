package http

import (
	"encoding/json"
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
)

type UserHandler struct{
	UserService *services.UserService
}


func NewUserHandler(userService *services.UserService) *UserHandler{
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) getCheckUsers(w http.ResponseWriter, r *http.Request){
	response := h.UserService.CheckAPI()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.AK_USERS

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if  user.Username =="" || user.Password =="" || user.Address == "" || user.RoleId == "" || user.DepartementId == "" || user.FullName == ""{
		http.Error(w, "Data Pengguna Tidak Lengkap", http.StatusBadRequest)
		return
	}
	// hashPassword := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// user.Password = string(hashPassword)
	response := h.UserService.CreateUser(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request){
	response := h.UserService.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}