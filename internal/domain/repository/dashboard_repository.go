package repository

import "gorm.io/gorm"

type DashboardRepository interface {
	CountTotalUsers() (int64, error)
	CountOnLeaveToday(date string) (int64, error)
	CountPresentToday(date string) (int64, error)
	CountAbsentToday(date, day string) (int64, error)
}

type dashboardRepository struct {
	DB *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) DashboardRepository {
	return &dashboardRepository{DB: db}
}

func (r *dashboardRepository) CountTotalUsers() (int64, error) {
	var count int64
	err := r.DB.Table("ak_users").Count(&count).Error
	return count, err
}

func (r *dashboardRepository) CountOnLeaveToday(date string) (int64, error) {
	var count int64

	err := r.DB.Table("ak_izins").
		Where("? BETWEEN start_date AND end_date AND status = '1'", date).
		Count(&count).Error
	return count, err
}

func (r *dashboardRepository) CountPresentToday(date string) (int64, error) {
	var count int64
	err := r.DB.Table("ak_kehadirans").Where("tanggal = ?", date).Count(&count).Error
	return count, err
}

func (r *dashboardRepository) CountAbsentToday(date, day string) (int64, error) {
	var count int64

	err := r.DB.Table("ak_schedules s").
		Where("s.day = ? AND s.is_active = 1", day).
		Where("NOT EXISTS (SELECT 1 FROM ak_kehadirans k WHERE k.user_uid = s.user_uid AND k.tanggal = ?)", date).
		Where("NOT EXISTS (SELECT 1 FROM ak_izins i WHERE i.user_uid = s.user_uid AND ? BETWEEN i.start_date AND i.end_date AND i.status = '1')", date).
		Count(&count).Error
	return count, err
}
