package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/gorilla/mux"
)

type RadiusHandler struct {
	RadiusService *services.RadiusService
}

func NewRadiusHandler(radiusService *services.RadiusService) *RadiusHandler {
	return &RadiusHandler{RadiusService: radiusService}

}

func (h *RadiusHandler) GetAllRadius(w http.ResponseWriter, r *http.Request) {
	response := h.RadiusService.GetAllRadius()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *RadiusHandler) CreateRadius(w http.ResponseWriter, r *http.Request) {
	var radius models.Ak_Radius
	if err := json.NewDecoder(r.Body).Decode(&radius); err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
	}

	response := h.RadiusService.CreateRadius(&radius)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *RadiusHandler) GetRadiusById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.RadiusService.GetRadiusById(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *RadiusHandler) UpdateRadius(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var radius models.Ak_Radius

	radius.RadiusId = int(num)

	response := h.RadiusService.UpdateRadius(&radius)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h *RadiusHandler) DeleteRadius(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.RadiusService.DeleteRadius(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
