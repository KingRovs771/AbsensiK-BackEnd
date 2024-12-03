package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type TipePotonganService struct {
	TipePotonganRepo *repository.TipePotonganRepository
}

func NewTipePotonganService(tipePotonganRepo *repository.TipePotonganRepository) *TipePotonganService {
	return &TipePotonganService{TipePotonganRepo: tipePotonganRepo}
}

func (s *TipePotonganService) GetAllTipePotongan() map[string]interface{} {
	TipePotongan, err := s.TipePotonganRepo.GetAllTipePotongan()

	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Ditemukan",
			"Error":   err.Error(),
		}
	}

	if len(TipePotongan) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Belum Ada Data Tipe Potongan",
		}
	}

	return map[string]interface{}{
		"Status":        "Success",
		"Message":       "Data Berhasil Ditemukan",
		"Tipe Potongan": TipePotongan,
	}
}

func (s *TipePotonganService) CreateTipePotongan(tipePotongan *models.Ak_TipePotongan) map[string]interface{} {
	if err := s.TipePotonganRepo.CreateTipePotongan(tipePotongan); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Menginput Data Tipe Potongan",
			"Error":   err.Error(),
		}
	}

	if tipePotongan.NamaPotongan == "" || tipePotongan.NilaiPotongan == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Tolong Lengkapi Data Potongan",
		}
	}

	return map[string]interface{}{
		"Status":        "Success",
		"Message":       "Data Berhasil Di Inputkan",
		"Tipe Potongan": tipePotongan,
	}
}

func (s *TipePotonganService) GetTipePotonganById(TipePotonganId int64) map[string]interface{} {
	tipePotongan, err := s.TipePotonganRepo.GetTipePotonganById(TipePotonganId)

	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Ditemukan, silakan check lagi",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":        "Success",
		"Message":       "Data Berhasil Ditemukan",
		"Tipe Potongan": tipePotongan,
	}
}

func (s *TipePotonganService) UpdateTipePotongan(TipePotongan *models.Ak_TipePotongan) map[string]interface{} {
	if err := s.TipePotonganRepo.UpdateTipePotongan(TipePotongan); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tipe Potongan Tidak Berhasil Di update",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":        "Success",
		"Message":       "Data Berhasil Di Update Silakan Check Data",
		"Tipe Potongan": TipePotongan,
	}
}

func (s *TipePotonganService) DeleteTipePotongan(TipePotonganId int64) map[string]interface{} {
	if err := s.TipePotonganRepo.DeleteTipePotongan(TipePotonganId); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Berhasil Di Update",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data Berhasil Di Hapus",
	}
}
