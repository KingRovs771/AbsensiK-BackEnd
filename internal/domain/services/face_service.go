package services

import (
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
