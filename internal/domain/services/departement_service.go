package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type DepartementService struct {
	DepartementsRepository *repository.DepartementsRepository
	SecretKey              string
}

func NewDepartementService(departementRepo *repository.DepartementsRepository, secretKey string) *DepartementService {
	return &DepartementService{DepartementsRepository: departementRepo}
}

func (s *DepartementService) CheckAPI() string {
	return `{"message":"berhasil"}`
}

func (s *DepartementService) GetDepartementsAll() map[string]interface{} {
	departements, err := s.DepartementsRepository.GetAllDepartements()
	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Error",
			"Error":   err.Error(),
		}

	}

	if len(departements) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Departements Not Found",
		}
	}

	return map[string]interface{}{
		"Status":      "Success",
		"Message":     "Berhasil Mendapatkan Data Departements",
		"Departments": departements,
	}
}

func (s *DepartementService) CreateDepartements(departement *models.Ak_Departments) map[string]interface{} {
	if err := s.DepartementsRepository.CreateDepartements(departement); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Menambahkan Departements Baru",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":      "Success",
		"Message":     "Berhasil Menambahkan Departements Baru",
		"Departemens": departement,
	}
}

func (s *DepartementService) GetDepartementsById(DepartementsId string) map[string]interface{} {
	departements, err := s.DepartementsRepository.GetDepartementsById(DepartementsId)

	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Departements Tidak Ditemukan",
			"Error":   err.Error(),
		}
	}
	return map[string]interface{}{
		"Status":       "Success",
		"Message":      "Berhasil Mendapatkan Departements",
		"Departements": departements,
	}
}

func (s *DepartementService) UpdateDepartements(departements *models.Ak_Departments) map[string]interface{} {
	if err := s.DepartementsRepository.UpdateDepartement(departements); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Memperbarui Departements",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":       "Success",
		"Message":      "Berhasil Memperbarui Departements",
		"Departements": departements,
	}
}

func (s *DepartementService) DeleteDepartements(DepartementsId string) map[string]interface{} {
	if err := s.DepartementsRepository.DeleteDepartement(DepartementsId); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Menghapus Departements",
			"Error":   err.Error(),
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Menghapus Departements",
	}
}
