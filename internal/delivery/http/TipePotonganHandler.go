package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
)

type TipePotonganHandler struct {
	TipePotonganService *services.TipePotonganService
}

func NewTipePotonganHandler(tipePotonganService *services.TipePotonganService) *TipePotonganHandler {
	return &TipePotonganHandler{TipePotonganService: tipePotonganService}
}

func (h *TipePotonganHandler) GetAllTipePotongan(w http.ResponseWriter, r *http.Request) {
	response := h.TipePotonganService.GetAllTipePotongan()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TipePotonganHandler) CreateTipePotongan(w http.ResponseWriter, r *http.Request) {
	var TipePotongan models.Ak_TipePotongan
	if err := json.NewDecoder(r.Body).Decode(&TipePotongan); err != nil {
		log.Println("Invalid Payload Request")
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
	}

	response := h.TipePotonganService.CreateTipePotongan(&TipePotongan)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *TipePotonganHandler) GetTipePotonganById(w http.ResponseWriter, r *http.Request) {

}

func (h *TipePotonganHandler) UpdateTipePotongan(w http.ResponseWriter, r *http.Request) {

}

func (h *TipePotonganHandler) DeleteTipePotongan(w http.ResponseWriter, r *http.Request) {

}
