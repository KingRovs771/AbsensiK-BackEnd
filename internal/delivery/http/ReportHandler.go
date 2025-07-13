package http

import (
	"encoding/json"
	"fmt"
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
		http.Error(w, `{"Status": "Error", "Message": "Tipe laporan tidak valid"}`, http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"Status": "Error", "Message": "Gagal menghasilkan laporan: %v"}`, err), http.StatusInternalServerError)
		return
	}
	message := "Laporan berhasil diambil"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Status":  "Success",
		"Message": message,
		"Data":    data,
	})
}
