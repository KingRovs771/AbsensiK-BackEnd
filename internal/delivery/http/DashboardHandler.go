package http

import (
	"encoding/json"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"net/http"
)

type DashboardHandler struct {
	DashboardService services.DashboardService
}

func NewDashboardHandler(service services.DashboardService) *DashboardHandler {
	return &DashboardHandler{DashboardService: service}
}

func (h *DashboardHandler) GetDashboardStats(w http.ResponseWriter, r *http.Request) {
	stats, err := h.DashboardService.GetDashboardStats()
	if err != nil {
		http.Error(w, "Gagal mengambil data statistik", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Status":  "Success",
		"Message": "Data statistik berhasil diambil",
		"Data":    stats,
	})
}
