package repository

import (
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type FaceRepository struct {
	DB *gorm.DB
}

func NewFaceRepository(db *gorm.DB) *FaceRepository {
	return &FaceRepository{DB: db}
}

func (r *FaceRepository) GetAllFaces() ([]models.Ak_Face, error) {
	var faces []models.Ak_Face

	if err := r.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_uid, full_name")
	}).Select("faces_id, user_uid, face_data").Find(&faces).Error; err != nil {
		fmt.Println("Error fetching face data with user:", err)
		return nil, err
	}
	return faces, nil
}

func (r *FaceRepository) SaveFaceData(face *models.Ak_Face) error {
	if err := r.DB.Create(face).Error; err != nil {
		fmt.Println("Error saving face data:", err)
		return err
	}
	return nil
}

func (r *FaceRepository) GetFaceByUserID(userUID string) (*models.Ak_Face, error) {
	var face models.Ak_Face
	if err := r.DB.Where("user_uid = ?", userUID).First(&face).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		fmt.Println("Error Fetching Data Face", err)
		return nil, err
	}
	return &face, nil
}

func (r *FaceRepository) DeleteFoto(UserUID string) error {
	return r.DB.Where("user_uid= ?", UserUID).Delete(&models.Ak_Face{}).Error
}

func (r *FaceRepository) GetFotoByID(id int) (*models.Ak_Face, error) {
	var face models.Ak_Face
	err := r.DB.Preload("User").First(&face, id).Error
	return &face, err
}

func (r *FaceRepository) UpdateFoto(face *models.Ak_Face) error {

	updateData := map[string]interface{}{
		"user_uid": face.UserUID,
	}

	if len(face.FaceData) > 0 {
		updateData["face_data"] = face.FaceData
	}

	return r.DB.Model(&models.Ak_Face{}).
		Where("faces_id = ?", face.FacesID).
		Updates(updateData).Error
}
