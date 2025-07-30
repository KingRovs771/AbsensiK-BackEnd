package services

import (
	"encoding/base64"
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type FaceService struct {
	FaceRepository *repository.FaceRepository
}

func NewFaceService(faceRepository *repository.FaceRepository) *FaceService {
	return &FaceService{FaceRepository: faceRepository}
}
func (s *FaceService) GetAllFaces() ([]models.Ak_Face, error) {
	return s.FaceRepository.GetAllFaces()
}

func (s *FaceService) UploadFace(face *models.Ak_Face) error {
	if face.UserUID == "" || len(face.FaceData) == 0 {
		fmt.Println("error : Data Wajah Tidak Valid")
		return fmt.Errorf("Invalid Face Data")
	}
	return s.FaceRepository.SaveFaceData(face)
}

func (s *FaceService) GetFaceByUserUID(userUID string) (*models.Ak_Face, error) {
	return s.FaceRepository.GetFaceByUserID(userUID)
}

func (s *FaceService) DeleteFotoUser(UserUID string) map[string]interface{} {
	if err := s.FaceRepository.DeleteFoto(UserUID); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Menghapus Pengguna",
			"Error":   err.Error(),
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Menghapus Pengguna",
	}
}

func (s *FaceService) GetFotoByID(id int) map[string]interface{} {
	// Panggil repository untuk mendapatkan data dari database
	face, err := s.FaceRepository.GetFotoByID(id)
	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data foto tidak ditemukan",
			"Error":   err.Error(),
		}
	}

	response := struct {
		FaceID   int64           `json:"face_id"`
		UserUID  string          `json:"user_uid"`
		FaceData string          `json:"face_data"` // Diubah menjadi string untuk Base64
		User     models.Ak_Users `json:"user"`
	}{
		FaceID:   face.FacesID,
		UserUID:  face.UserUID,
		FaceData: base64.StdEncoding.EncodeToString(face.FaceData),
		User:     face.User,
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data foto berhasil ditemukan",
		"Data":    response,
	}
}

func (s *FaceService) UpdateFoto(face *models.Ak_Face) map[string]interface{} {
	// Panggil repository untuk melakukan update data di database
	if err := s.FaceRepository.UpdateFoto(face); err != nil {
		// Jika terjadi error, kembalikan respons error
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal memperbarui data foto di database",
			"Error":   err.Error(),
		}
	}

	// Jika berhasil, kembalikan respons sukses
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data foto berhasil diperbarui",
	}
}
