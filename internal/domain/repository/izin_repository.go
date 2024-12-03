package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type IzinRepository struct {
	DB *gorm.DB
}

func NewIzinRepository(db *gorm.DB) *IzinRepository {
	return &IzinRepository{DB: db}
}

func (r *IzinRepository) GetAllIzin() ([]models.Ak_Izin, error) {
	var izin []models.Ak_Izin
	if err := r.DB.Find(&izin).Error; err != nil {
		return nil, err
	}

	return izin, nil
}

func (r *IzinRepository) CreateIzin(izin *models.Ak_Izin) error {
	if err := r.DB.Create(izin).Error; err != nil {
		return err
	}

	return nil
}

func (r *IzinRepository) GetIzinById(IzinId int64) (*models.Ak_Izin, error) {
	var izin models.Ak_Izin

	if err := r.DB.First(&izin, IzinId).Error; err != nil {
		return nil, err
	}
	return &izin, nil
}

func (r *IzinRepository) UpdateIzin(izin *models.Ak_Izin) error {
	return r.DB.Save(izin).Error
}

func (r *IzinRepository) DeleteIzin(IzinId int64) error {
	return r.DB.Delete(&models.Ak_Izin{}, IzinId).Error
}
