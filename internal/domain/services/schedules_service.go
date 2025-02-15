package services

import (
	"errors"
	"log"
	"time"

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
		"Status":  "Success",
		"Message": "Data Schedules Ditemukan",
		"Data":    schedules,
	}
}

func (s *SchedulesService) CreateSchedules(startTime, endTime, userId, day string, isActive int) map[string]interface{} {

	if err := s.ValidateTime(startTime); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Invalid Start Time Format",
			"Error":   err.Error(),
		}
	}
	if err := s.ValidateTime(startTime); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Invalid Start Time Format",
			"Error":   err.Error(),
		}
	}

	schedules := &models.Ak_Schedules{
		UserUID:   userId,
		StartTime: startTime,
		EndTime:   endTime,
		Day:       day,
		IsActive:  isActive,
	}

	if err := s.SchedulesRepository.CreateSchedules(schedules); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Berhasil Di simpan",
			"Error":   err.Error(),
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data Schedules Berhasil Disimpan",
		"Data":    schedules,
	}
}

func (s *SchedulesService) ValidateTime(timeStr string) error {
	layout := "15:04"
	_, err := time.Parse(layout, timeStr)
	if err != nil {
		log.Println("Invalid Time Format:", err)
		return errors.New("Invalid time format")
	}
	return nil
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

func (s *SchedulesService) DeleteSchedules(ScheduleId int64) map[string]interface{} {
	if err := s.SchedulesRepository.DeleteSchedules(ScheduleId); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Berhasil Dihapus",
			"Error":   err.Error(),
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data Berhasil Dihapus",
		"Id":      ScheduleId,
	}
}
