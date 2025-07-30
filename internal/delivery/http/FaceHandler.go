package http

import (
	"encoding/json"
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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
func (h *FaceHandler) GetFotoByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID foto tidak ditemukan di parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID tidak valid", http.StatusBadRequest)
		return
	}

	response := h.FaceService.GetFotoByID(id)

	w.Header().Set("Content-Type", "application/json")

	if response["Status"] == "Error" {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(response)
}
func (h *FaceHandler) UpdateFoto(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Face ID is missing in parameters", http.StatusBadRequest)
		return
	}
	faceID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Face ID", http.StatusBadRequest)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Could not parse multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Dapatkan user_uid dari form value
	userUID := r.FormValue("user_uid")
	if userUID == "" {
		http.Error(w, "User UID is required", http.StatusBadRequest)
		return
	}

	faceDataToUpdate := models.Ak_Face{
		FacesID: int64(faceID),
		UserUID: userUID,
	}

	file, _, err := r.FormFile("face_data")
	if err != nil {
		if err != http.ErrMissingFile {
			http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading the file: "+err.Error(), http.StatusInternalServerError)
			return
		}
		faceDataToUpdate.FaceData = fileBytes
	}

	response := h.FaceService.UpdateFoto(&faceDataToUpdate)

	w.Header().Set("Content-Type", "application/json")
	if response["Status"] == "Error" {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(response)
}
