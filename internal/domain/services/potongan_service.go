package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type PotonganService struct {
	PotonganRepository *repository.PotonganRepository
}

func NewPotonganService(potonganRepository *repository.PotonganRepository) *PotonganService {
	return &PotonganService{PotonganRepository: potonganRepository}
}

func (s *PotonganService) GetAllPotongan() map[string]interface{} {
	potongan, err := s.PotonganRepository.GetAllPotongan()

	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Error",
			"Error":   err.Error(),
		}
	}

	if len(potongan) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Potongan Not Found",
		}
	}

	return map[string]interface{}{
		"Status":   "Success",
		"Message":  "Data Potongan Berhasil di dapatkan",
		"Potongan": potongan,
	}
}

func (s *PotonganService) CreatePotongan(potongan *models.Ak_Potongan) map[string]interface{} {
	if err := s.PotonganRepository.CreatePotongan(potongan); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Error",
			"Error":   err.Error(),
		}
	}

	if potongan.UserId == "" || potongan.Month == "" || potongan.TipePotongan == "" {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Belum Lengkap",
		}
	}

	return map[string]interface{}{
		"Status":   "Success",
		"Message":  "Data Berhasil Disimpan",
		"Potongan": potongan,
	}
}

func (s *PotonganService) GetPotonganById(PotonganId int64) map[string]interface{} {
	potongan, err := s.PotonganRepository.GetPotonganById(PotonganId)

	if err != nil {
		return map[string]interface{}{
			"Status ": "Error",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":   "Success",
		"Message":  "Berhasil Mendapatkan Potongan",
		"Potongan": potongan,
	}
}

func (s *PotonganService) UpdatePotongan(Potongan *models.Ak_Potongan) map[string]interface{} {
	if err := s.PotonganRepository.UpdatePotongan(Potongan); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Berhasil Di Update",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":   "Error",
		"Message":  "Data Berhasil Di Update",
		"Potongan": Potongan,
	}
}

func (s *PotonganService) DeletePotongan(PotonganId int64) map[string]interface{} {
	if err := s.PotonganRepository.DeletePotongan(PotonganId); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Potongan Berhasil Di Delete",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Error",
		"Message": "Berhasil Menghapus Data Potongan",
		"Data":    PotonganId,
	}
}
