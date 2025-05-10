package services

import "github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"

type KehadiranService struct {
	KehadiranRepository *repository.KehadiranRepository
	SecretKey           string
}

func NewKehadiranService(kehadiranRepo *repository.KehadiranRepository, secretKey string) *KehadiranService {
	return &KehadiranService{
		KehadiranRepository: kehadiranRepo,
		SecretKey:           secretKey,
	}
}
