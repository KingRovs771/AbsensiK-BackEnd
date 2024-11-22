package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/golang-jwt/jwt/v5"
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

func (h *UserHandler)GetProfile(w http.ResponseWriter, r *http.Request){
	authHeader := r.Header.Get("Authorization")
	if authHeader == ""{
		log.Println("Authorization Header Required")
		http.Error(w, "Authorization Header Required", http.StatusUnauthorized)
		return
	}
	tokenString := strings.Split(authHeader, "Bearer ")[1]
	claims := jwt.MapClaims{}
	token,err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error){
		return []byte(h.UserService.SecretKey), nil
	})

	if err != nil || !token.Valid {
		log.Println("Invalid Token", err)
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}
	UserId := int64(claims["user_id"].(float64))
	user, err := h.UserService.GetProfile(UserId)

	if err != nil{
		log.Println("User Not Foud :", err)
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}
	response := map[string]interface{}{
		"Status" : "Success",
		"Message" : "User Di Temukan",
		"User" : user,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}