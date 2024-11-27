package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/gorilla/mux"
)

type RoleHandler struct {
	RoleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{RoleService: roleService}
}

func (h *RoleHandler) GetAllRole(w http.ResponseWriter, r *http.Request) {
	response := h.RoleService.GetAllService()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *RoleHandler) CreateRole(w http.ResponseWriter, r *http.Request) {
	var roles models.Ak_Roles
	if err := json.NewDecoder(r.Body).Decode(&roles); err != nil {
		log.Println("Invalid Payload Request")
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
	}

	response := h.RoleService.CreateService(&roles)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *RoleHandler) GetRoleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response := h.RoleService.GetRolesById(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func (h *RoleHandler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var roles models.Ak_Roles

	if err := json.NewDecoder(r.Body).Decode(&roles); err != nil {
		log.Println("Invalid Request Payload")
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
	}

	roles.RoleId = id
	response := h.RoleService.UpdateRole(&roles)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *RoleHandler) DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response := h.RoleService.DeleteRole(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
