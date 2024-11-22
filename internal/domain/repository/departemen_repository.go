package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type DepartementRepo struct {
	DB *gorm.DB
}

func NewDepartementsRepository(db *gorm.DB) *DepartementRepo {
	return &DepartementRepo{DB: db}
}
func (r *DepartementRepo) getAllDepartements() ([]models.Ak_Departments, error) {
	var departemen []models.Ak_Departments

	if err := r.DB.Find(&departemen).Error; err != nil {
		return nil, err
	}
	return departemen, nil
}

func (r *DepartementRepo) getDepartementsById(DepartementId string) (*models.Ak_Departments, error) {
	var departemen models.Ak_Departments
	if err := r.DB.First(&departemen, DepartementId).Error; err != nil {
		return nil, err
	}
	return &departemen, nil
}

func (r *DepartementRepo) createDepartements(departement *models.Ak_Departments) error {
	return r.DB.Create(departement).Error
}

func (r *DepartementRepo) updateDepartements(departement *models.Ak_Departments) error {
	return r.DB.Save(departement).Error
}

func (r *DepartementRepo) deleteDepartements(DepartementId string) error {
	return r.DB.Delete(&models.Ak_Departments{}, DepartementId).Error
}
