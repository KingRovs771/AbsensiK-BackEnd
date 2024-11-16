package http

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
)

type UserHandler struct{
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler{
	return &UserHandler{UserService: userService}
}