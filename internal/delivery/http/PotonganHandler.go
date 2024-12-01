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

type PotonganHandler struct {
	PotonganService *services.PotonganService
}

func NewPotonganHandler(potonganHandler *services.PotonganService) *PotonganHandler {
	return &PotonganHandler{PotonganService: potonganHandler}
}

func (h *PotonganHandler) GetAllPotongan(w http.ResponseWriter, r *http.Request) {
	response := h.PotonganService.GetAllPotongan()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PotonganHandler) CreatePotongan(w http.ResponseWriter, r *http.Request) {
	var potongan models.Ak_Potongan
	if err := json.NewDecoder(r.Body).Decode(&potongan); err != nil {
		log.Println("Invalid Payload Request")
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
	}

	response := h.PotonganService.CreatePotongan(&potongan)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PotonganHandler) GetPotonganById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("Error")
		return
	}
	response := h.PotonganService.GetPotonganById(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PotonganHandler) UpdatePotongan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Println("Error Not Found Id")
		return
	}

	var potongan models.Ak_Potongan

	potongan.PotonganId = num

	response := h.PotonganService.UpdatePotongan(&potongan)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *PotonganHandler) DeletePotongan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Println("Invalid Id Potongan")
		return
	}

	response := h.PotonganService.DeletePotongan(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
