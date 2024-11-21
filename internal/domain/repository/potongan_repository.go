package repository

import "gorm.io/gorm"

type PotonganRepository struct {
	DB *gorm.DB
}

func NewPotonganRepository(db *gorm.DB) *PotonganRepository {
	return &PotonganRepository{DB:db}
}

