package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type IzinService struct {
	IzinRepository *repository.IzinRepository
}

func NewIzinService(izinRepository *repository.IzinRepository) *IzinService {
	return &IzinService{IzinRepository: izinRepository}
}

func (s *IzinService) GetAllIzin() map[string]interface{} {
	izin, err := s.IzinRepository.GetAllIzin()
	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Izin Tidak Ditemukan",
			"Error":   err.Error(),
		}
	}

	if len(izin) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Izin Belum Ada",
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data Izin Ditemukan",
		"Izin":    izin,
	}
}

func (s *IzinService) CreateIzin(izin *models.Ak_Izin) map[string]interface{} {
	if err := s.IzinRepository.CreateIzin(izin); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Izin Tidak Berhasil Di Inputkan",
			"Error":   err.Error(),
		}
	}

	if izin.UserId == "" || izin.IzinType == "" || izin.Alasan == "" {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Izin Harus Dilengkapi",
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data Berhasil Disimpan",
		"izin":    izin,
	}
}

func (s *IzinService) GetIzinById(izinId int64) map[string]interface{} {
	izin, err := s.IzinRepository.GetIzinById(izinId)

	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Ditemukan",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Succes",
		"Message": "Data Berhasil Ditemukan",
		"izin":    izin,
	}
}

func (s *IzinService) UpdateIzin(izin *models.Ak_Izin) map[string]interface{} {
	if err := s.IzinRepository.UpdateIzin(izin); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Berhasil Di Update",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data Berhasil Di Update",
		"Izin":    izin,
	}
}

func (s *IzinService) DeleteIzin(IzinId int64) map[string]interface{} {
	if err := s.IzinRepository.DeleteIzin(IzinId); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Tidak Menghapus Data Izin",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Menghapus Data Izin",
		"Izin":    IzinId,
	}
}
