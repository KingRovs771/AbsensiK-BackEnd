package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type RadiusRepository struct {
	DB *gorm.DB
}

func NewRadiusRepository(db *gorm.DB) *RadiusRepository {
	return &RadiusRepository{DB: db}
}
func (r *RadiusRepository) GetAllRadius() ([]models.Ak_Radius, error) {
	var radius []models.Ak_Radius

	if err := r.DB.Find(&radius).Error; err != nil {
		return nil, err
	}
	return radius, nil
}

func (r *RadiusRepository) GetRadiusById(RadiusId int64) (*models.Ak_Radius, error) {
	var radius models.Ak_Radius
	if err := r.DB.First(&radius, RadiusId).Error; err != nil {
		return nil, err
	}
	return &radius, nil
}

func (r *RadiusRepository) CreateRadius(radius *models.Ak_Radius) error {
	return r.DB.Create(radius).Error
}

func (r *RadiusRepository) UpdateRadius(Radius *models.Ak_Radius) error {
	return r.DB.Save(Radius).Error
}

func (r *RadiusRepository) DeleteRadius(RadiusId int64) error {
	return r.DB.Delete(&models.Ak_Radius{}, RadiusId).Error
}
