package repository

import "gorm.io/gorm"

type FaceRepository struct {
	DB *gorm.DB
}

func NewFaceRepository(db *gorm.DB) *FaceRepository {
	return &FaceRepository{DB: db}
}
