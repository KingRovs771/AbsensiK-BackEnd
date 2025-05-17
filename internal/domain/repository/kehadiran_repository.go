package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
	"time"
)

type KehadiranRepository struct {
	DB *gorm.DB
}

func NewKehadiranRepository(db *gorm.DB) *KehadiranRepository {
	return &KehadiranRepository{DB: db}
}

func (r *KehadiranRepository) GetKehadiran() ([]models.Ak_Kehadiran, error) {
	var kehadiran []models.Ak_Kehadiran

	if err := r.DB.Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_uid, full_name")
	}).Preload("Schedules ").Find(&kehadiran).Error; err != nil {
		return nil, err
	}
	return kehadiran, nil
}

func (r *KehadiranRepository) GetKehadiranById(KehadiranId int64) (*models.Ak_Kehadiran, error) {
	var kehadiran models.Ak_Kehadiran
	if err := r.DB.First(&kehadiran, KehadiranId).Error; err != nil {
		return nil, err
	}
	return &kehadiran, nil
}

func (r *KehadiranRepository) InsertCheckIn(userID string, latitude, longitude float64, photo string, checkInTime time.Time) error {
	kehadiran := models.Ak_Kehadiran{
		UserId:    userID,
		Latitude:  latitude,
		Longitude: longitude,
		Photo:     photo,
		TimeIn:    &checkInTime,
	}
	return r.DB.Create(&kehadiran).Error
}

func (r *KehadiranRepository) GetLatestCheckIn(userID string, attendance *models.Ak_Kehadiran) error {
	return r.DB.Where("user_id = ? AND time_out IS NULL", userID).First(attendance).Error
}

func (r *KehadiranRepository) InsertCheckOut(kehadiran *models.Ak_Kehadiran) error {
	return r.DB.Save(&kehadiran).Error
}
