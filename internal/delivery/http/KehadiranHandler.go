package http

import (
	"encoding/json"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"net/http"
	"strconv"
)

type KehadiranHandler struct {
	KehadiranService *services.KehadiranService
}

func NewKehadiranHandler(kehadiranService *services.KehadiranService) *KehadiranHandler {
	return &KehadiranHandler{KehadiranService: kehadiranService}
}

func (h *KehadiranHandler) GetKehadiran(w http.ResponseWriter, r *http.Request) {
	response := h.KehadiranService.GetAllKehadiran()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}

func (h *KehadiranHandler) CreateKehadiran(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	userID, _ := strconv.ParseUint(r.FormValue("user_id"), 10, 64)
	latitude, _ := strconv.ParseFloat(r.FormValue("latitude"), 64)
	longitude, _ := strconv.ParseFloat(r.FormValue("longitude"), 64)
	photo := r.FormValue("photo")

	response := h.KehadiranService.InsertCheckIn(string(userID), latitude, longitude, photo)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}
