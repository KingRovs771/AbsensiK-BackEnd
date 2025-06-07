package repository

import (
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type PotonganRepository struct {
	DB *gorm.DB
}

func NewPotonganRepository(db *gorm.DB) *PotonganRepository {
	return &PotonganRepository{DB: db}
}

func (r *PotonganRepository) GetAllPotongan(month string, year int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	err := r.DB.Table("ak_potongans").
		Select("ak_potongans.potongan_id, ak_potongans.user_uid, ak_potongans.month, ak_potongans.year, ak_users.full_name, ak_potongans.tipe_potongan").
		Joins("JOIN ak_users ON ak_potongans.user_uid = ak_users.user_uid").
		Joins("JOIN ak_tipe_potongans ON ak_potongans.tipe_potongan = ak_tipe_potongans.tipe_potongan_id").
		Where("ak_potongans.month = ? AND ak_potongans.year = ?", month, year).
		Find(&results).Error

	if err != nil {
		fmt.Println("Error fetching salary data:", err)
		return nil, err
	}
	return results, nil
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
