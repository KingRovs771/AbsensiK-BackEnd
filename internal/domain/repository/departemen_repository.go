package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type DepartementRepo struct {
	DB *gorm.DB
}

func NewDepartementRepository(db *gorm.DB) *DepartementRepo{
	return &DepartementRepo{DB : db}
}
func (r *DepartementRepo) getAllDepartement()([]models.Ak_Departement, error){
	  var departemen []models.Ak_Departement

	  if err := r.DB.Find(&departemen).Error; err != nil {
		return nil, err
	  }
	  return departemen, nil
}

func (r * DepartementRepo) getDepartementById(DepartementId string)(*models.Ak_Departement, error){
	var departemen models.Ak_Departement
	if err := r.DB.First(&departemen, DepartementId).Error; err != nil {
		return nil, err
	}
	return &departemen, nil
}

func (r *DepartementRepo) createDepartement(departement *models.Ak_Departement) error{
	return r.DB.Create(departement).Error
}

func (r *DepartementRepo) updateDepartement(departement *models.Ak_Departement) error{
	return r.DB.Save(departement).Error
}

func (r *DepartementRepo) deleteDepartement(DepartementId string) error{
	return r.DB.Delete(&models.Ak_Departement{}, DepartementId ).Error
}