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

func (r *UserRepository) GetAllUsers() ([]GetAllUserTable, error) {
	var users []GetAllUserTable
	if err := r.DB.Table("absensi_karyawan.ak_users").
		Select("absensi_karyawan.ak_users.user_id, absensi_karyawan.ak_users.full_name, absensi_karyawan.ak_users.departement_id, absensi_karyawan.ak_users.role_id, absensi_karyawan.ak_departments.name_departments, absensi_karyawan.ak_roles.name_role").
		Joins("JOIN absensi_karyawan.ak_roles ON absensi_karyawan.ak_users.role_id = absensi_karyawan.ak_roles.role_id").
		Joins("JOIN absensi_karyawan.ak_departments ON absensi_karyawan.ak_users.departement_id = absensi_karyawan.ak_departments.departments_id").
		Scan(&users).Error; err != nil {
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

func (r *UserRepository) DeleteUser(UserId int64) error {
	return r.DB.Delete(&models.Ak_Users{}, UserId).Error
}
