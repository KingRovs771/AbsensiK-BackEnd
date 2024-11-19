package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepository: userRepo}
}

func(s *UserService) CheckAPI() string{
	return `{"message":"berhasil"}`
}

func (s *UserService) CreateUser(user *models.AK_USERS)map[string]interface{}{
	if err := s.UserRepository.CreateUser(user); err != nil {
		return map[string]interface{}{
			"Status" : "Error",
			"Message" : "Gagal Membuat User Baru",
			"Error" : err.Error(),
		}
		
	}
	return map[string]interface{}{
		"Status" : "Success",
		"Message" : "Berhasil Membuat Pengguna Baru",
		"Users" : user,
	}
}

func (s *UserService) GetAllUsers() map[string]interface{}{
	users, err := s.UserRepository.GetAllUsers()
	if err != nil {
		return map[string]interface{}{
			"Status" : "Error",
			"Message" : "Gagal Membuat User Baru",
			"Error" : err.Error(),
		}
	}

	if len(users) == 0{
		return map[string]interface{}{
			"Status" : "Error",
			"Message" : "Belum ada Data",
		}
	}

	return map[string]interface{}{
		"Status" : "Success",
		"Message" : "Berhasil Mendapatkan Semua Users",
		"Users" : users,
	}
}