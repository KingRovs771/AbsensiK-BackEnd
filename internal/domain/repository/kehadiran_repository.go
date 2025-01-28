package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type KehadiranRepository struct {
	DB *gorm.DB
}

func NewKehadiranRepository(db *gorm.DB) *KehadiranRepository {
	return &KehadiranRepository{DB: db}
}

func (r *KehadiranRepository) GetKehadiran() ([]models.Ak_Kehadiran, error) {
	var kehadiran []models.Ak_Kehadiran

	if err := r.DB.Find(&kehadiran).Error; err != nil {
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

func (r *KehadiranRepository) UpdateKehadiran(kehadiran *models.Ak_Kehadiran) error {
	return r.DB.Save(kehadiran).Error
}

func (r *KehadiranRepository) DeleteKehadiran(KehadiranId int64) error {
	return r.DB.Delete(&models.Ak_Kehadiran{}, KehadiranId).Error
}
