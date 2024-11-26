package http

import (
	"encoding/json"
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
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

}
