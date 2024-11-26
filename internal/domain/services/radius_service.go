package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type RadiusService struct {
	RadiusRepository *repository.RadiusRepository
}

func NewRadiusService(radiusRepository *repository.RadiusRepository) *RadiusService {
	return &RadiusService{RadiusRepository: radiusRepository}
}

func (s *RadiusService) GetAllRadius() map[string]interface{} {
	radius, err := s.RadiusRepository.GetAllRadius()
	if err != nil {
		return map[string]interface{}{
			"Status": "Error",
			"Error":  err.Error(),
		}
	}

	if len(radius) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Radius Tidak Ada",
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Mendapatkan Data Radius",
		"Radius":  radius,
	}
}

func (s *RadiusService) CreateRadius(radius *models.Ak_Radius) map[string]interface{} {
	if radius.NameLocation == "" || radius.Latitude == 0 || radius.Longitude == 0 || radius.Radius == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Silakan Isi Kelengkapan Data Radius, Tidak boleh ada yang kosong",
		}
	}

	if err := s.RadiusRepository.CreateRadius(radius); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Tidak Dapat Membuat Radius Baru",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Sukses Membuat Radius Baru",
		"Radius":  radius,
	}
}

func (s *RadiusService) GetRadiusById(RadiusId int64) map[string]interface{} {
	radius, err := s.RadiusRepository.GetRadiusById(RadiusId)

	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Pengguna tidak ditemukan",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Pengguna ditemukan",
		"Radius":  radius,
	}
}
func (s *RadiusService) UpdateRadius(radius *models.Ak_Radius) map[string]interface{} {
	if err := s.RadiusRepository.UpdateRadius(radius); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Radius Gagal Terupdate",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Error",
		"Message": "Data Radius Berhasil Terupdate",
		"Radius":  radius,
	}
}

func (s *RadiusService) DeleteRadius(RadiusId int64) map[string]interface{} {
	if err := s.RadiusRepository.DeleteRadius(RadiusId); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Menghapus Data Radius",
			"Data":    RadiusId,
		}
	}

	return map[string]interface{}{
		"Status":    "Success",
		"Message":   "Berhasil Menghapus Data Radius",
		"Radius Id": RadiusId,
	}
}
