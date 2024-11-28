package http

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/gorilla/mux"
)

type SchedulesHandler struct {
	SchedulesService *services.SchedulesService
}

func NewSchedulesHandler(schedulesService *services.SchedulesService) *SchedulesHandler {
	return &SchedulesHandler{SchedulesService: schedulesService}
}

func (h *SchedulesHandler) GetAllSchedules(w http.ResponseWriter, r *http.Request) {
	response := h.SchedulesService.GetAllSchedules()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SchedulesHandler) CreateSchedules(w http.ResponseWriter, r *http.Request) {
	var schdules models.Ak_Schedules
	if err := json.NewDecoder(r.Body).Decode(&schdules); err != nil {
		log.Println("Invalid Payload Request")
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
	}

	response := h.SchedulesService.CreateSchedules(&schdules)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SchedulesHandler) GetSchedulesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := h.SchedulesService.GetSchedulesById(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SchedulesHandler) UpdateSchedules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var schdules models.Ak_Schedules

	schdules.SchedulesId = num

	if err := json.NewDecoder(r.Body).Decode(schdules); err != nil {
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
		return
	}

	response := h.SchedulesService.UpdateSchedules(&schdules)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *SchedulesHandler) DeleteSchedules(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := h.SchedulesService.DeleteSchedules(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
