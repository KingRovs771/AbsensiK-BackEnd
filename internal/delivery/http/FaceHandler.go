package http

import (
	"encoding/json"
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type FaceHandler struct {
	FaceService *services.FaceService
}

func NewFaceHandler(faceService *services.FaceService) *FaceHandler {
	return &FaceHandler{FaceService: faceService}
}

func (h *FaceHandler) GetAllFacesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Metode harus GET", http.StatusMethodNotAllowed)
		return
	}

	faces, err := h.FaceService.GetAllFaces()
	if err != nil {
		http.Error(w, "Gagal mengambil data wajah", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(faces)
}

func (h *FaceHandler) UploadFaceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode harus POST", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20) // 🔹 Maksimum 10MB
	if err != nil {
		fmt.Println("Error parsing multipart form:", err)
		http.Error(w, "Gagal membaca form-data", http.StatusBadRequest)
		return
	}

	userUID := r.FormValue("user_uid")
	fmt.Println("Received userUID:", userUID)

	if userUID == "" {
		http.Error(w, "User UID tidak terkirim", http.StatusBadRequest)
		return
	}

	// 🔹 Validasi file foto yang diunggah
	file, _, err := r.FormFile("face_data")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 🔹 Baca file sebagai byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Gagal membaca file", http.StatusInternalServerError)
		return
	}

	// 🔹 Simpan data ke database
	faceData := &models.Ak_Face{
		UserUID:  userUID,
		FaceData: fileBytes,
	}

	err = h.FaceService.UploadFace(faceData)
	if err != nil {
		http.Error(w, "Gagal menyimpan data wajah", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Face data uploaded successfully",
	})
}

func (h *FaceHandler) GetFaceByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userUID := r.URL.Query().Get("user_uid")
	if userUID == "" {
		http.Error(w, "User UID harus disertakan", http.StatusBadRequest)
		return
	}

	faceData, err := h.FaceService.GetFaceByUserUID(userUID)
	if err != nil {
		http.Error(w, "Gagal mengambil data wajah", http.StatusInternalServerError)
		return
	}
	if faceData == nil {
		http.Error(w, "Data wajah tidak ditemukan", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(faceData.FaceData)
}

func (h *FaceHandler) DeleteFoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response := h.FaceService.DeleteFotoUser(id)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(response)
}
