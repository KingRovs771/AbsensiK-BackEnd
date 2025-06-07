package services

import "github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"

type FaceService struct {
	FaceRepository *repository.FaceRepository
}

func NewFaceService(faceRepository *repository.FaceRepository) *FaceService {
	return &FaceService{FaceRepository: faceRepository}
}
