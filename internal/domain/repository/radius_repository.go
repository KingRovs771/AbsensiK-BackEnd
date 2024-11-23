package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type RadiusRepository struct {
	DB *gorm.DB
}

func NewRadiusRepository(db *gorm.DB) *RadiusRepository{
	return &RadiusRepository{DB : db}
}
func (r *RadiusRepository) getAllRadius()([]models.Ak_Radius, error){
	  var radius []models.Ak_Radius

	  if err := r.DB.Find(&radius).Error; err != nil {
		return nil, err
	  }
	  return radius, nil
}

func (r * RadiusRepository) getRadiusById(RadiusId string)(*models.Ak_Radius, error){
	var radius models.Ak_Radius
	if err := r.DB.First(&radius, RadiusId ).Error; err != nil {
		return nil, err
	}
	return &radius, nil
}

func (r *RadiusRepository) createRadius(radius *models.Ak_Radius) error{
	return r.DB.Create(radius).Error
}

func (r *RadiusRepository) updateRadius(Radius *models.Ak_Radius) error{
	return r.DB.Save(Radius).Error
}

func (r *RadiusRepository) deleteRadius(RadiusId string) error{
	return r.DB.Delete(&models.Ak_Radius{}, RadiusId ).Error
}