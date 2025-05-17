package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
	"time"
)

type KehadiranService struct {
	KehadiranRepository *repository.KehadiranRepository
	SecretKey           string
}

func NewKehadiranService(Repo *repository.KehadiranRepository) *KehadiranService {
	return &KehadiranService{KehadiranRepository: Repo}
}

func (s *KehadiranService) CheckAPI() string {
	return `message : Berhasil Menyambung`
}

func (s *KehadiranService) GetAllKehadiran() map[string]interface{} {
	kehadiran, err := s.KehadiranRepository.GetKehadiran()
	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Tidak Ditemukan",
			"Error":   err.Error(),
		}
	}

	if len(kehadiran) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Data Kehadiran Belum Ada",
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Data Kehadiran Ditemukan",
		"Data":    kehadiran,
	}
}

func (s *KehadiranService) InsertCheckIn(userID string, latitude, longitude float64, photo string) map[string]interface{} {
	checkInTime := time.Now()

	err := s.KehadiranRepository.InsertCheckIn(userID, latitude, longitude, photo, checkInTime)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Failed to insert check-in",
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Check-in successful",
		"data": map[string]interface{}{
			"user_id":   userID,
			"latitude":  latitude,
			"longitude": longitude,
			"photo":     photo,
			"check_in":  checkInTime,
		},
	}
}

func (s *KehadiranService) InsertCheckOut(userID string, latitude, longitude float64, photo string) map[string]interface{} {
	checkOutTime := time.Now()
	attendance := models.Ak_Kehadiran{}

	err := s.KehadiranRepository.GetLatestCheckIn(userID, &attendance)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "No check-in record found",
		}
	}

	attendance.Latitude = latitude
	attendance.Longitude = longitude
	attendance.Photo = photo
	attendance.TimeOut = &checkOutTime

	err = s.KehadiranRepository.InsertCheckOut(&attendance)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Failed to update check-out",
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Check-out successful",
		"data": map[string]interface{}{
			"user_id":   userID,
			"latitude":  latitude,
			"longitude": longitude,
			"photo":     photo,
			"check_out": checkOutTime,
		},
	}
}
