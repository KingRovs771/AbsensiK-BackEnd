package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type ReportRepository interface {
	GetSalaryReport(month, year string) ([]models.SalaryReport, error)
	GetAttendanceReport(month, year string) ([]models.AttendanceReport, error)
	GetLeaveReport(month, year string) ([]models.LeaveReport, error)
	GetAllUsers() ([]models.Ak_Users, error)
}

type reportRepository struct {
	DB *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{DB: db}
}

func (r *reportRepository) GetSalaryReport(month, year string) ([]models.SalaryReport, error) {
	var results []models.SalaryReport
	err := r.DB.Table("ak_salaries s").
		Select("s.user_uid, u.full_name, s.month, s.year, s.total_gaji, s.total_potongan").
		Joins("left join ak_users u on u.user_uid = s.user_uid").
		Where("s.month = ? AND s.year = ?", month, year).
		Find(&results).Error
	return results, err
}

func (r *reportRepository) GetAttendanceReport(month, year string) ([]models.AttendanceReport, error) {
	var results []models.AttendanceReport
	// Menggunakan LIKE untuk mencocokkan bulan dan tahun dari tanggal
	datePattern := year + "-" + month + "-%"
	err := r.DB.Table("ak_kehadirans k").
		Select("k.user_uid, u.full_name, k.tanggal, k.time_in, k.time_out").
		Joins("left join ak_users u on u.user_uid = k.user_uid").
		Where("k.tanggal LIKE ?", datePattern).
		Order("k.tanggal asc, k.time_in asc").
		Find(&results).Error
	return results, err
}

func (r *reportRepository) GetLeaveReport(month, year string) ([]models.LeaveReport, error) {
	var results []models.LeaveReport
	datePattern := year + "-" + month + "-%"
	err := r.DB.Table("ak_izins i").
		Select("i.user_uid, u.full_name, i.izin_type as permit_type, i.start_date, i.end_date, i.alasan, i.status").
		Joins("left join ak_users u on u.user_uid = i.user_uid").
		Where("i.start_date LIKE ?", datePattern).
		Find(&results).Error
	return results, err
}

func (r *reportRepository) GetAllUsers() ([]models.Ak_Users, error) {
	var users []models.Ak_Users
	err := r.DB.Find(&users).Error
	return users, err
}
