package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler{
	return &AuthHandler{
		AuthService: authService,
	}
}

type LoginRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err !=nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	token, err := h.AuthService.Authenticate(req.Username, req.Password)
	if err != nil{
		log.Println("Error Authenricating User : ", err)
		http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
		return
	} 

	log.Println("User Authenticated Successfully :", req.Username)

	response := map[string]string{
		"Status" : "Success",
		"Token" : token,
		"Username" : req.Username, 
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}