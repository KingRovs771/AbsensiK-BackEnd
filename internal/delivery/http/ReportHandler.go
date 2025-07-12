package http

import (
	"encoding/json"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"net/http"
)

type ReportHandler struct {
	ReportService services.ReportService
}

func NewReportHandler(service services.ReportService) *ReportHandler {
	return &ReportHandler{ReportService: service}
}

func (h *ReportHandler) GenerateReport(w http.ResponseWriter, r *http.Request) {
	reportType := r.URL.Query().Get("type")
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	var data interface{}
	var err error

	switch reportType {
	case "salary":
		data, err = h.ReportService.GetSalaryReport(month, year)
	case "attendance":
		data, err = h.ReportService.GetAttendanceReport(month, year)
	case "leave":
		data, err = h.ReportService.GetLeaveReport(month, year)
	case "users":
		data, err = h.ReportService.GetAllUsers()
	default:
		http.Error(w, "Tipe laporan tidak valid", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Gagal menghasilkan laporan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   data,
	})
}
