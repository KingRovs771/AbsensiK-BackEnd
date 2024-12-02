package services

import "github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"

type TipePotonganService struct {
	TipePotonganRepo *repository.TipePotonganRepository
}

func NewTipePotonganService(tipePotonganRepo *repository.TipePotonganRepository) *TipePotonganService {
	return &TipePotonganService{TipePotonganRepo: tipePotonganRepo}
}
