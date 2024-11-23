package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type DepartementsRepository struct {
	DB *gorm.DB
}

func NewDepartementsRepository(db *gorm.DB) *DepartementsRepository {
	return &DepartementsRepository{DB: db}
}

func (r *DepartementsRepository) CreateDepartements(departement *models.Ak_Departments) error {
	if err := r.DB.Create(departement).Error; err != nil {
		return err
	}
	return nil
}

func (r *DepartementsRepository) GetAllDepartements() ([]models.Ak_Departments, error) {
	var departement []models.Ak_Departments
	if err := r.DB.Find(&departement).Error; err != nil {
		return nil, err
	}
	return departement, nil
}

func (r *DepartementsRepository) GetDepartementsById(DepartementsId string) (*models.Ak_Departments, error) {
	var departement models.Ak_Departments
	if err := r.DB.First(&departement, DepartementsId).Error; err != nil {
		return nil, err
	}
	return &departement, nil
}

func (r *DepartementsRepository) UpdateDepartement(departements *models.Ak_Departments) error {
	return r.DB.Save(departements).Error
}

func (r *DepartementsRepository) DeleteDepartement(DepartementsId string) error {
	return r.DB.Delete(&models.Ak_Departments{}, DepartementsId).Error
}
