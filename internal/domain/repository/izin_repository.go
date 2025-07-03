package repository

import (
	"log"
	"time"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type IzinRepository struct {
	DB *gorm.DB
}

func NewIzinRepository(db *gorm.DB) *IzinRepository {
	return &IzinRepository{DB: db}
}

func (r *IzinRepository) GetAllIzin() ([]models.Ak_Izin, error) {
	var izin []models.Ak_Izin
	if err := r.DB.Preload("Users", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_uid, full_name")
	}).Find(&izin).Error; err != nil {
		return nil, err
	}

	return izin, nil
}

func (r *IzinRepository) CreateIzin(izin *models.Ak_Izin) error {
	if err := r.DB.Create(izin).Error; err != nil {
		return err
	}

	return nil
}

func (r *IzinRepository) GetIzinById(IzinId int64) (*models.Ak_Izin, error) {
	var izin models.Ak_Izin

	if err := r.DB.First(&izin, IzinId).Error; err != nil {
		return nil, err
	}
	return &izin, nil
}

func (r *IzinRepository) UpdateIzin(izin *models.Ak_Izin) error {
	return r.DB.Save(izin).Error
}

func (r *IzinRepository) DeleteIzin(IzinId int64) error {
	return r.DB.Delete(&models.Ak_Izin{}, IzinId).Error
}

func (r *IzinRepository) ApproveIzin(IzinId int64, ApproveBy string) error {
	updateData := map[string]interface{}{
		"status":       1,
		"approve_by":   ApproveBy,
		"approve_date": time.Now(),
	}
	err := r.DB.Model(&models.Ak_Izin{}).Where("izin_id = ?", IzinId).Updates(updateData).Error
	if err != nil {
		log.Println("Error Approving Izin : ", err)
		return err
	}
	return nil
}

func (r *IzinRepository) GetPermitsByUserUID(userUID string) ([]models.Ak_Izin, error) {
	var permits []models.Ak_Izin

	// Mencari semua data di tabel 'ak_izins' yang cocok dengan user_uid.
	// Kita juga mengurutkannya dari yang terbaru (berdasarkan timestamp).
	err := r.DB.Where("user_uid = ?", userUID).Order("created_at desc").Find(&permits).Error
	if err != nil {
		return nil, err
	}

	return permits, nil
}
