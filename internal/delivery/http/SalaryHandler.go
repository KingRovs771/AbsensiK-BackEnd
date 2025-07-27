package http

import (
	"encoding/json"
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"net/http"
	"strings"
	"time"
)

type SalaryHandler struct {
	SalaryService *services.SalaryService
	AuthService   *services.AuthService
}

type SalaryRequest struct {
	UserUID string `json:"user_uid"`
	Month   string `json:"month"`
	Year    int64  `json:"year"`
}

// Struct untuk error helper
type appError struct {
	Message string
	Code    int
}

func (e *appError) Error() string {
	return e.Message
}

func NewSalaryHandler(salaryService *services.SalaryService, authService *services.AuthService) *SalaryHandler {
	return &SalaryHandler{SalaryService: salaryService, AuthService: authService}
}

func (h *SalaryHandler) GetSalary(w http.ResponseWriter, r *http.Request) {
	response := h.SalaryService.CheckAPI()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SalaryHandler) GenerateSalary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON dari frontend
	var req SalaryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to parse JSON data", http.StatusBadRequest)
		return
	}

	fmt.Println("Received userUID:", req.UserUID)
	// Panggil service untuk menghitung gaji
	response := h.SalaryService.GenerateSalary(req.UserUID, req.Month, req.Year)

	// Set response headers dan kirim JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}

func (h *SalaryHandler) GetSalariesByMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	month := r.URL.Query().Get("month")
	currentYear := time.Now().Year() // 🔹 Otomatis ambil tahun saat ini

	if month == "" {
		http.Error(w, "Month parameter is required", http.StatusBadRequest)
		return
	}
	response := h.SalaryService.GetSalariesByMonth(month, currentYear)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}

func (h *SalaryHandler) GetSalariesByMonthAndName(w http.ResponseWriter, r *http.Request) {
	month := r.URL.Query().Get("month")
	year := time.Now().Year()

	responseData := h.SalaryService.GetSalariesByMonthAndName(month, year)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(responseData)
}

// Update
func (h *SalaryHandler) getUserFromRequest(r *http.Request) (*models.Ak_Users, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, &appError{"Authorization header required", http.StatusUnauthorized}
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return nil, &appError{"Invalid token format", http.StatusUnauthorized}
	}
	user, err := h.AuthService.GetUserFromToken(tokenString)
	if err != nil {
		return nil, &appError{"Invalid token", http.StatusUnauthorized}
	}
	return user, nil
}

func (h *SalaryHandler) GetLatestPayslip(w http.ResponseWriter, r *http.Request) {
	// 1. Panggil helper untuk otentikasi
	user, err := h.getUserFromRequest(r)
	if err != nil {
		if appErr, ok := err.(*appError); ok {
			http.Error(w, appErr.Message, appErr.Code)
		} else {
			http.Error(w, "An internal error occurred", http.StatusInternalServerError)
		}
		return
	}

	// 2. Lanjutkan dengan logika bisnis
	payslipData, err := h.SalaryService.GetLatestPayslip(user.UserUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "success", "data": payslipData})
}
