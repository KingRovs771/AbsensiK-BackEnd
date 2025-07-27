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
	var TipePotongan models.Ak_TipePotongans
	if err := json.NewDecoder(r.Body).Decode(&TipePotongan); err != nil {
		log.Println("Invalid Payload Request")
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
	}

	response := h.TipePotonganService.CreateTipePotongan(&TipePotongan)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}

func (h *TipePotonganHandler) GetTipePotonganById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["tipe_potongan_id"]
	num, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.TipePotonganService.GetTipePotonganById(num)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}

func (h *TipePotonganHandler) UpdateTipePotongan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["tipe_potongan_id"]
	num, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tipePotongan models.Ak_TipePotongans

	tipePotongan.TipePotonganId = int(num)
	var tipePotonganPayload models.Ak_TipePotongans
	if err := json.NewDecoder(r.Body).Decode(&tipePotonganPayload); err != nil {
		http.Error(w, "Invalid Request Payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	response := h.TipePotonganService.UpdateTipePotongan(&tipePotongan)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}

func (h *TipePotonganHandler) DeleteTipePotongan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["tipe_potongan_id"]
	num, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.TipePotonganService.DeleteTipePotongan(num)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}
