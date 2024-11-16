package services

import "github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(db *repository.UserRepository) *UserService {
	return &UserService{UserRepository: db}
}