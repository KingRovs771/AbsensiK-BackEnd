package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type SchedulesService struct {
	SchedulesRepository *repository.SchedulesRepository
}

func NewSchedulesService(schedulesRepository *repository.SchedulesRepository) *SchedulesService {
	return &SchedulesService{SchedulesRepository: schedulesRepository}
}

func (s *SchedulesService) GetAllSchedules() map[string]interface{} {
	schedules, err := s.SchedulesRepository.GetAllSchedules()
	if err != nil {
		return map[string]interface{}{
			"Status": "Error",
			"Error":  err.Error(),
		}
	}
	if len(schedules) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Ditemukan",
		}
	}
	return map[string]interface{}{
		"Status":    "Success",
		"Message":   "Data Schedules Ditemukan",
		"Schedules": schedules,
	}
}

func (s *SchedulesService) CreateSchedules(schedules *models.Ak_Schedules) map[string]interface{} {
	if schedules.SchedulesId == 0 || schedules.Day == "" {
		return map[string]interface{}{
			"Status ": "Error",
			"Message": "Data Belum Di isi Lengkap",
		}
	}

	if err := s.SchedulesRepository.CreateSchedules(schedules); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Berhasil Di simpan",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":    "Success",
		"Message":   "Data Schedules Berhasil Disimpan",
		"Schedules": schedules,
	}
}

func (s *SchedulesService) GetSchedulesById(SchedulesId int64) map[string]interface{} {
	schdules, err := s.SchedulesRepository.GetSchedulesById(SchedulesId)

	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Ditemukan",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":    "Success",
		"Message":   "Data ditemukan",
		"Schedules": schdules,
	}
}

func (s *SchedulesService) UpdateSchedules(schedules *models.Ak_Schedules) map[string]interface{} {
	if err := s.SchedulesRepository.UpdateSchedules(schedules); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Berhasil Di Update",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":    "Success",
		"Message":   "Data Berhasil Di Update",
		"Schedules": schedules,
	}
}

func (s *SchedulesService) DeleteSchedules(SchedulesId int64) map[string]interface{} {
	if err := s.SchedulesRepository.DeleteSchedules(SchedulesId); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Berhasil Dihapus",
			"Error":   err.Error(),
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data Berhasil Dihapus",
		"Id":      SchedulesId,
	}
}
