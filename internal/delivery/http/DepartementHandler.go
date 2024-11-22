package http

import "github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"

type DepartementHandler struct {
	DepartementService *services.DepartementService
}

func NewDepartementHandler(departementService *services.DepartementService) *DepartementHandler {
	return &DepartementHandler{DepartementService: departementService}
}
