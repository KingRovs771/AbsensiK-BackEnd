package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type GetAllUserTable struct {
	UserId          int64  `json:"user_id"`
	FullName        string `json:"full_name"`
	DepartementId   string `json:"departement_id"`
	RoleId          string `json:"role_id"`
	NameDepartments string `json:"name_departments"`
	NameRole        string `json:"name_role"`
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.Ak_Users) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetAllUsers() ([]models.Ak_Users, error) {
	var users []models.Ak_Users
	if err := r.DB.Preload("Department", func(db *gorm.DB) *gorm.DB {
		return db.Select("departments_id, name_departments")
	}).Preload("Role", func(db *gorm.DB) *gorm.DB {
		return db.Select("role_id, name_role")
	}).Select("user_uid, email, address, username, gender, full_name, dailyrate, phone,                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   departments_id, role_id").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserById(UserId int64) (*models.Ak_Users, error) {
	var users models.Ak_Users
	if err := r.DB.Preload("Ak_Roles").Preload("Ak_Departments").First(&users, UserId).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) GetUserByIdUpdate(UserId int64) (*models.Ak_Users, error) {
	var user models.Ak_Users
	if err := r.DB.First(&user, UserId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.Ak_Users) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) DeleteUser(UserUID string) error {
	return r.DB.Where("user_uid= ?", UserUID).Delete(&models.Ak_Users{}).Error
}

func (r *UserRepository) SearchEmployeeByName(name string) ([]models.Ak_Users, error) {
	var Ak_users []models.Ak_Users

	err := r.DB.Where("full_name LIKE ?", "%"+name+"%").Find(&Ak_users).Error
	return Ak_users, err
}
