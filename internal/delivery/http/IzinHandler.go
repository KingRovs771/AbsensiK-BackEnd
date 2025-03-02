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

type IzinHandler struct {
	IzinService *services.IzinService
}

func NewIzinHandler(izinService *services.IzinService) *IzinHandler {
	return &IzinHandler{IzinService: izinService}
}
func (h *IzinHandler) GetAllIzin(w http.ResponseWriter, r *http.Request) {
	response := h.IzinService.GetAllIzin()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *IzinHandler) CreateIzin(w http.ResponseWriter, r *http.Request) {
	var izin models.Ak_Izin
	if err := json.NewDecoder(r.Body).Decode(&izin); err != nil {
		log.Println("Invalid Payload Request")
		http.Error(w, "Invalid Payload Request", http.StatusBadRequest)
	}

	response := h.IzinService.CreateIzin(&izin)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *IzinHandler) GetIzinById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.IzinService.GetIzinById(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *IzinHandler) UpdateIzin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var izin models.Ak_Izin

	izin.IzinId = num

	response := h.IzinService.UpdateIzin(&izin)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *IzinHandler) DeleteIzin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := h.IzinService.DeleteIzin(num)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *IzinHandler) ApproveIzin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	izinIdStr, ok := vars["id"]

	if !ok || izinIdStr == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid Id")
		return
	}

	izinId, err := strconv.ParseInt(izinIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Id Parsing Id", http.StatusBadRequest)
		log.Println("Error Parsing Id")
		return
	}

	approveBy := r.Header.Get("X-Full-Name")
	if approveBy == "" {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Error : Tidak ada Approve")
		return
	}

	err = h.IzinService.ApprovingIzin(izinId, approveBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Izin Approve Success",
		},
	)
}
