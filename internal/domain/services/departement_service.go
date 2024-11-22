package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type DepartementService struct {
	DepartementRepo *repository.DepartementRepo
	SecretKey       string
}

func NewDepartementService(departementRepo *repository.DepartementRepo, secretKey string) *DepartementService {
	return &DepartementService{DepartementRepo: departementRepo}
}

func (s *DepartementService) CheckAPI() string {
	return `{"message":"berhasil"}`
}

func (s *DepartementService) CreateDepartements(departement *models.Ak_Departments) map[string]interface{} {
	return s.DepartementRepo.
}
