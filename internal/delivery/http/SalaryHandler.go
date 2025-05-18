package http

import (
	"encoding/json"
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"net/http"
	"time"
)

type SalaryHandler struct {
	SalaryService *services.SalaryService
}

type SalaryRequest struct {
	UserUID string `json:"user_uid"`
	Month   string `json:"month"`
	Year    int64  `json:"year"`
}

func NewSalaryHandler(salaryService *services.SalaryService) *SalaryHandler {
	return &SalaryHandler{SalaryService: salaryService}
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

	fmt.Println("Handler Debug - Month:", month, "Year:", currentYear) // 🔹 Debugging backend
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
	year := time.Now().Year() // 🔹 Tahun otomatis

	fmt.Println("Handler Debug - Month:", month, "Year:", year) // Debugging

	responseData := h.SalaryService.GetSalariesByMonthAndName(month, year)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(responseData)
}
