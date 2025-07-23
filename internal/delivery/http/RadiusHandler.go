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
	// 1. Decode the request into a temporary map first, not directly to the model.
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// 2. Extract and validate each field manually.
	nameLocation, _ := payload["name_location"].(string)
	latStr, _ := payload["latitude"].(string)
	lonStr, _ := payload["longitude"].(string)
	// Radius can be float64 from JSON, so we handle that.
	radiusFloat, _ := payload["radius"].(float64)

	if nameLocation == "" || latStr == "" || lonStr == "" {
		http.Error(w, "Name, latitude, and longitude are required", http.StatusBadRequest)
		return
	}

	// 3. Convert string values to the correct numeric types.
	latitude, errLat := strconv.ParseFloat(latStr, 64)
	longitude, errLon := strconv.ParseFloat(lonStr, 64)

	if errLat != nil || errLon != nil {
		http.Error(w, "Invalid format for latitude or longitude", http.StatusBadRequest)
		return
	}

	// 4. Create the final model struct with the correctly typed data.
	radiusModel := &models.Ak_Radius{
		NameLocation: nameLocation,
		Latitude:     latitude,
		Longitude:    longitude,
		Radius:       int(radiusFloat), // Convert float64 to int
	}

	// 5. Pass the correctly formed model to the service.
	response := h.RadiusService.CreateRadius(radiusModel) // Assuming your service accepts the model

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
		http.Error(w, "Gagal parse ID", http.StatusBadRequest)
		return
	}

	type RadiusUpdateRequest struct {
		NameLocation string `json:"name_location"`
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
		Radius       int    `json:"radius"`
	}

	var requestPayload RadiusUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid Request Payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	lat, errLat := strconv.ParseFloat(requestPayload.Latitude, 64)
	lon, errLon := strconv.ParseFloat(requestPayload.Longitude, 64)

	// Jika konversi gagal, kirim error.
	if errLat != nil || errLon != nil {
		http.Error(w, "Latitude atau Longitude bukan angka yang valid", http.StatusBadRequest)
		return
	}

	radiusToUpdate := models.Ak_Radius{
		RadiusId:     num,
		NameLocation: requestPayload.NameLocation,
		Latitude:     lat, // Gunakan nilai float64 yang sudah dikonversi
		Longitude:    lon, // Gunakan nilai float64 yang sudah dikonversi
		Radius:       requestPayload.Radius,
	}

	// 4. Lanjutkan proses ke service layer dengan data yang sudah bersih.
	response := h.RadiusService.UpdateRadius(&radiusToUpdate)
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}
