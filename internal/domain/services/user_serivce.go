package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
	SecretKey      string
}

func NewUserService(userRepo *repository.UserRepository, secretKey string) *UserService {
	return &UserService{UserRepository: userRepo,
		SecretKey: secretKey}
}

func (s *UserService) CreateUser(user *models.Ak_Users) map[string]interface{} {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return map[string]interface{}{
			"Status":  "error",
			"Message": "Gagal enkripsi Password",
			"error":   err.Error(),
		}
	}
	user.Password = string(hasedPassword)
	if err := s.UserRepository.CreateUser(user); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Membuat User Baru",
			"Error":   err.Error(),
		}

	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Membuat Pengguna Baru",
		"Data":    user,
	}
}

func (s *UserService) GetAllUsers() map[string]interface{} {
	users, err := s.UserRepository.GetAllUsers()
	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Error",
			"Error":   err.Error(),
		}
	}

	if len(users) == 0 {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Belum ada Data",
		}
	}

	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Mendapatkan Semua Users",
		"Data":    users,
	}
}

func (s *UserService) GetProfile(UserId int64) (*models.Ak_Users, error) {
	return s.UserRepository.GetUserById(UserId)
}

func (s *UserService) GetUsersById(UserId int64) map[string]interface{} {
	user, err := s.UserRepository.GetUserByIdUpdate(UserId)
	if err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Pengguna Tidak Ditemukan",
			"Error":   err.Error(),
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Mendapatkan Pengguna",
		"Data":    user,
	}
}

func (s *UserService) UpdateUser(user *models.Ak_Users) map[string]interface{} {
	if err := s.UserRepository.UpdateUser(user); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Memperbarui Pengguna",
			"Error":   err.Error(),
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Memperbarui Pengguna",
		"Data":    user,
	}
}

func (s *UserService) DeleteUser(UserId int64) map[string]interface{} {
	if err := s.UserRepository.DeleteUser(UserId); err != nil {
		return map[string]interface{}{
			"Status":  "Error",
			"Message": "Gagal Menghapus Pengguna",
			"Error":   err.Error(),
		}
	}
	return map[string]interface{}{
		"Status":  "Success",
		"Message": "Berhasil Menghapus Pengguna",
	}
}

func (s *UserService) SearchEmployeeByName(name string) map[string]interface{} {
	employees, err := s.UserRepository.SearchEmployeeByName(name)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve employees",
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Employees retrieved successfully",
		"data":    employees,
	}
}
