package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/gorilla/mux"
)

type DepartementHandler struct {
	DepartementService *services.DepartementService
}

func NewDepartementHandler(departementService *services.DepartementService) *DepartementHandler {
	return &DepartementHandler{DepartementService: departementService}
}

func (h *DepartementHandler) GetAllDepartements(w http.ResponseWriter, r *http.Request) {
	response := h.DepartementService.GetDepartementsAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *DepartementHandler) CreateDepartements(w http.ResponseWriter, r *http.Request) {
	var departement models.Ak_Departments
	if err := json.NewDecoder(r.Body).Decode(&departement); err != nil {
		log.Println("Invalid Request Payload :", err)
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	response := h.DepartementService.CreateDepartements(&departement)
	w.Header().Set("Content-TYpe", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *DepartementHandler) GetDepartementsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response := h.DepartementService.GetDepartementsById(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *DepartementHandler) UpdateDepartement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var departement models.Ak_Departments

	if err := json.NewDecoder(r.Body).Decode(&departement); err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	departement.DepartmentsId = id

	response := h.DepartementService.UpdateDepartements(&departement)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *DepartementHandler) DeleteDepartement(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response := h.DepartementService.DeleteDepartements(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
