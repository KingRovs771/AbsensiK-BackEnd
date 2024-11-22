package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) CreateRole(role *models.Ak_Roles) error {
	return r.DB.Create(role).Error
}

func (r *RoleRepository) GetAllRole() ([]models.Ak_Roles, error) {
	var role []models.Ak_Roles
	if err := r.DB.Find(&role).Error; err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepository) GetRoleById(RoleId string) (*models.Ak_Roles, error) {
	var role models.Ak_Roles
	if err := r.DB.First(&role, RoleId).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) UpdateRole(role *models.Ak_Roles) error {
	return r.DB.Save(role).Error
}

func (r *RoleRepository) DeleteRole(RoleId string) error {
	return r.DB.Delete(&models.Ak_Roles{}, RoleId).Error
}
