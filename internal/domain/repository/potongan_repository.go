package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type PotonganRepository struct {
	DB *gorm.DB
}

func NewPotonganRepository(db *gorm.DB) *PotonganRepository {
	return &PotonganRepository{DB: db}
}

func (r *PotonganRepository) GetAllPotongan() ([]models.Ak_Potongan, error) {
	var potongan []models.Ak_Potongan
	if err := r.DB.Find(&potongan).Error; err != nil {
		return nil, err
	}
	return potongan, nil
}

func (r *PotonganRepository) CreatePotongan(potongan *models.Ak_Potongan) error {
	if err := r.DB.Create(potongan).Error; err != nil {
		return err
	}
	return nil
}

func (r *PotonganRepository) GetPotonganById(PotonganId int64) (*models.Ak_Potongan, error) {
	var potongan models.Ak_Potongan
	if err := r.DB.First(&potongan, PotonganId).Error; err != nil {
		return nil, err
	}
	return &potongan, nil
}

func (r *PotonganRepository) UpdatePotongan(potongan *models.Ak_Potongan) error {
	return r.DB.Save(potongan).Error
}

func (r *PotonganRepository) DeletePotongan(PotonganId int64) error {
	return r.DB.Delete(&models.Ak_Potongan{}, PotonganId).Error
}
