package http

import (
	"encoding/json"
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type IzinHandler struct {
	IzinService *services.IzinService
	AuthService *services.AuthService
}

func NewIzinHandler(izinService *services.IzinService, authService *services.AuthService) *IzinHandler {
	return &IzinHandler{IzinService: izinService, AuthService: authService}
}
func (h *IzinHandler) GetAllIzin(w http.ResponseWriter, r *http.Request) {
	response := h.IzinService.GetAllIzin()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *IzinHandler) CreateIzin(w http.ResponseWriter, r *http.Request) {
	// 1. Set max memory untuk form
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
		return
	}

	// 2. Baca semua field teks dari form
	var izin models.Ak_Izin

	izin.UserUID = r.FormValue("user_uid")
	izin.IzinType = r.FormValue("izin_type")
	izin.Alasan = r.FormValue("alasan")
	izin.StartDate = r.FormValue("start_date")
	izin.EndDate = r.FormValue("end_date")
	izin.Status = 0

	// 3. Handle file upload (jika ada)
	file, _, err := r.FormFile("foto")
	if err != nil {
		// Jika error bukan karena file tidak ada, berarti ini error sebenarnya
		if err != http.ErrMissingFile {
			log.Println("Error retrieving file from form-data:", err)
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		// Jika file tidak ada (untuk form Izin biasa), biarkan field Photo nil
	} else {
		defer file.Close()

		// === PERUBAHAN UTAMA DI SINI ===
		// Baca seluruh isi file sebagai array byte ([]byte)
		photoBytes, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Unable to read file content", http.StatusInternalServerError)
			return
		}

		// Masukkan data byte ke dalam struct
		izin.Foto = photoBytes
	}

	// 4. Panggil service dengan data yang sudah di-parse
	response := h.IzinService.CreateIzin(&izin)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
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
	fmt.Println("Received izinIdStr:", izinIdStr)

	izinId, err := strconv.ParseInt(izinIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Id Parsing Id", http.StatusBadRequest)
		log.Println("Error Parsing Id")
		return
	}

	approveBy := r.Header.Get("X-Full-Name")
	if approveBy == "" {
		http.Error(w, "Header X-Full-Name tidak ditemukan", http.StatusBadRequest)
		log.Println("Error: Tidak ada Approve")
		return
	}

	err = h.IzinService.ApprovingIzin(izinId, approveBy)
	if err != nil {
		log.Println("Error approving izin:", err) // Debugging error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Full-Name")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Izin Approve Success",
		},
	)
}

func (h *IzinHandler) getUserFromRequest(r *http.Request) (*models.Ak_Users, error) {
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
func (h *IzinHandler) GetUserPermitHistory(w http.ResponseWriter, r *http.Request) {

	// 1. Panggil helper untuk otentikasi
	user, err := h.getUserFromRequest(r)
	if err != nil {
		appErr := err.(*appError)
		http.Error(w, appErr.Message, appErr.Code)
		return
	}

	// 2. Lanjutkan dengan logika Anda
	permits, err := h.IzinService.GetPermitsByUserUID(user.UserUID)
	if err != nil {
		http.Error(w, "Could not fetch permit history", http.StatusInternalServerError)
		return
	}

	if permits == nil {
		permits = []models.Ak_Izin{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "success", "data": permits})
}
