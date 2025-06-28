package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	RoleId   string `json:"role_id"`
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error Decoding Credentials:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	tokenString, err := h.AuthService.Authenticate(req.Username, req.Password)
	if err != nil {
		log.Println("Error Authenticating User : ", err)
		http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
		return
	}

	log.Println("User Authenticated Successfully :", req.Username)

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	json.NewEncoder(w).Encode(map[string]string{
		"Status":   "Success",
		"Token":    tokenString,
		"Username": req.Username,
		"Message":  "Login Success",
	})
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})
	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Logout Successfully",
		})
}

func (h *AuthHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := h.AuthService.GetUserFromToken(cookie.Value)
	if err != nil || user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) ProfileMobile(w http.ResponseWriter, r *http.Request) {
	// 1. Ambil token dari header Authorization
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header required", http.StatusUnauthorized)
		return
	}

	// Format header adalah "Bearer <token>"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader { // Jika tidak ada prefix "Bearer "
		http.Error(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	// 2. Gunakan service untuk mendapatkan user dari token
	user, err := h.AuthService.GetUserFromToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// 3. Kirim data user sebagai respons JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   user, // Mengirim seluruh objek user
	})
}
