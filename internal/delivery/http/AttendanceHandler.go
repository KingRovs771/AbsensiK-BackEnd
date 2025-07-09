package http

import (
	"encoding/json"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"net/http"
	"strings"
)

type AttendanceHandler struct {
	AttendanceService services.AttendanceService
	AuthService       *services.AuthService
}

func NewAttendanceHandler(attendanceService services.AttendanceService, authService *services.AuthService) *AttendanceHandler {
	return &AttendanceHandler{AttendanceService: attendanceService, AuthService: authService}
}

func (h *AttendanceHandler) getUserFromRequest(r *http.Request) (*models.Ak_Users, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, &appError{"Authorization header required", http.StatusUnauthorized}
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return nil, &appError{"Invalid token format", http.StatusUnauthorized}
	}
	user, err := h.AuthService.GetUserFromToken(tokenString)
	if err != nil {
		return nil, &appError{"Invalid token", http.StatusUnauthorized}
	}
	return user, nil
}

func (h *AttendanceHandler) GetAttendanceData(w http.ResponseWriter, r *http.Request) {
	user, err := h.getUserFromRequest(r)
	if err != nil {
		if appErr, ok := err.(*appError); ok {
			http.Error(w, appErr.Message, appErr.Code)
		} else {
			http.Error(w, "An internal error occurred", http.StatusInternalServerError)
		}
		return
	}

	data, err := h.AttendanceService.GetAttendancePageData(user.UserUID)
	if err != nil {
		http.Error(w, "Gagal mengambil data absensi", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "success", "data": data})
}

func (h *AttendanceHandler) ClockIn(w http.ResponseWriter, r *http.Request) {
	user, err := h.getUserFromRequest(r)
	if err != nil {
		if appErr, ok := err.(*appError); ok {
			http.Error(w, appErr.Message, appErr.Code)
		} else {
			http.Error(w, "An internal error occurred", http.StatusInternalServerError)
		}
		return
	}

	r.ParseMultipartForm(10 << 20)
	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")
	file, _, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Foto bukti wajib diunggah", http.StatusBadRequest)
		return
	}
	defer file.Close()

	message, err := h.AttendanceService.PerformClockIn(user.UserUID, file, latitude, longitude)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": message})
}

func (h *AttendanceHandler) ClockOut(w http.ResponseWriter, r *http.Request) {
	user, err := h.getUserFromRequest(r)
	if err != nil {
		if appErr, ok := err.(*appError); ok {
			http.Error(w, appErr.Message, appErr.Code)
		} else {
			http.Error(w, "An internal error occurred", http.StatusInternalServerError)
		}
		return
	}

	var payload struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}
	json.NewDecoder(r.Body).Decode(&payload)

	message, err := h.AttendanceService.PerformClockOut(user.UserUID, payload.Latitude, payload.Longitude)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": message})
}
