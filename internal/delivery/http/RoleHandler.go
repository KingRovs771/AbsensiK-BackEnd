package http

import "github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"

type RoleHandler struct {
	RoleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) *RoleHandler {
	return &RoleHandler{RoleService: roleService}
}
