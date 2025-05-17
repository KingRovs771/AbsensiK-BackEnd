package http

import (
	"encoding/json"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"net/http"
	"strconv"
)

type SalaryHandler struct {
	SalaryService *services.SalaryService
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

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	userID := r.FormValue("user_id")
	month := r.FormValue("month")
	year, _ := strconv.ParseInt(r.FormValue("year"), 10, 64)

	response := h.SalaryService.GenerateSalary(userID, month, year)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SalaryHandler) GetSalariesByCurrentMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	response := h.SalaryService.GetSalariesByCurrentMonth()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
