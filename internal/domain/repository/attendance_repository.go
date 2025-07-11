package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	DB *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{DB: db}
}

func (r *AttendanceRepository) GetTodaySchedule(userUID string, day string) (*models.Ak_Schedules, error) {
	var schedule models.Ak_Schedules
	err := r.DB.Where("user_uid = ? AND day = ? AND is_active = ?", userUID, day, 1).First(&schedule).Error
	return &schedule, err
}

func (r *AttendanceRepository) GetOfficeLocation() (*models.Ak_Radius, error) {
	var location models.Ak_Radius
	err := r.DB.First(&location).Error
	return &location, err
}

func (r *AttendanceRepository) GetTodayAttendance(userUID string, date string) (*models.Ak_Kehadiran, error) {
	var attendance models.Ak_Kehadiran
	err := r.DB.Where("user_uid = ? AND tanggal = ?", userUID, date).First(&attendance).Error
	return &attendance, err
}

func (r *AttendanceRepository) CreateClockIn(attendance *models.Ak_Kehadiran) error {
	return r.DB.Create(attendance).Error
}

func (r *AttendanceRepository) UpdateClockOut(userUID string, date string, timeOut string) error {
	return r.DB.Model(&models.Ak_Kehadiran{}).Where("user_uid = ? AND tanggal = ?", userUID, date).Update("time_out", timeOut).Error
}

func (r *AttendanceRepository) GetAllAttendancesWithUser(date string) ([]models.Ak_Kehadiran, error) {
	var results []models.Ak_Kehadiran

	err := r.DB.Table("ak_kehadiran").
		Select("ak_kehadiran.kehadiran_id, ak_kehadiran.user_uid, ak_users.full_name, ak_kehadiran.tanggal, ak_kehadiran.time_in, ak_kehadiran.time_out").
		Joins("left join ak_users on ak_users.user_uid = ak_kehadiran.user_uid").
		Where("ak_kehadiran.tanggal = ?", date).
		Order("ak_kehadiran.time_in asc").
		Find(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil
}
