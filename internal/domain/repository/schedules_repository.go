package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type SchedulesRepository struct {
	DB *gorm.DB
}

func NewSchedulesRepository(db *gorm.DB) *SchedulesRepository {
	return &SchedulesRepository{DB: db}
}

func (r *SchedulesRepository) GetAllSchedules() ([]models.Ak_Schedules, error) {
	var schedules []models.Ak_Schedules

	if err := r.DB.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil

}

func (r *SchedulesRepository) CreateSchedules(schedules *models.Ak_Schedules) error {
	return r.DB.Create(schedules).Error
}

func (r *SchedulesRepository) GetSchedulesById(SchedulesId int64) (*models.Ak_Schedules, error) {
	var schedules models.Ak_Schedules
	if err := r.DB.First(&schedules, SchedulesId).Error; err != nil {
		return nil, err
	}
	return &schedules, nil
}

func (r *SchedulesRepository) UpdateSchedules(schedules *models.Ak_Schedules) error {
	return r.DB.Save(schedules).Error
}
func (r *SchedulesRepository) DeleteSchedules(SchedulesId int64) error {
	return r.DB.Delete(&models.Ak_Schedules{}, SchedulesId).Error
}
