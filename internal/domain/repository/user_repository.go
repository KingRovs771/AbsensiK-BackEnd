package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.AK_USERS) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAllUsers() ([]models.AK_USERS, error) {
	var users []models.AK_USERS
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserById(UserId int64) (*models.AK_USERS, error) {
	var users models.AK_USERS
	if err := r.DB.Preload("Ak_Roles").Preload("Ak_Departement").First(&users, UserId).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) GetUserByIdUpdate(UserId int64) (*models.AK_USERS, error) {
	var user models.AK_USERS
	if err := r.DB.First(&user, UserId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.AK_USERS) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(UserId int64) error {
	return r.DB.Delete(&models.AK_USERS{}, UserId).Error
}
