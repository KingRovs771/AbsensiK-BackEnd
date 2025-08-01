package repository

import (
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
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

func normalizeMonth(monthInput string) (string, error) {
	// Pertama, coba konversi input menjadi angka
	if monthNum, err := strconv.Atoi(monthInput); err == nil {
		// Jika berhasil (input adalah "7" atau "07"), konversi angka ke nama bulan
		if monthNum < 1 || monthNum > 12 {
			return "", fmt.Errorf("nomor bulan tidak valid: %d", monthNum)
		}
		return time.Month(monthNum).String(), nil
	}

	// Jika bukan angka, anggap sebagai nama bulan. Standarkan ke format "Titlecase".
	// Ini akan mengubah "july", "JULY", atau "July" menjadi "July".
	lowerMonth := strings.ToLower(monthInput)
	var monthMap = map[string]string{
		"january": "January", "february": "February", "march": "March",
		"april": "April", "may": "May", "june": "June",
		"july": "July", "august": "August", "september": "September",
		"october": "October", "november": "November", "december": "December",
	}

	monthName, ok := monthMap[lowerMonth]
	if !ok {
		return "", fmt.Errorf("nama bulan tidak valid: %s", monthInput)
	}
	return monthName, nil
}
func (r *reportRepository) GetSalaryReport(month, year string) ([]models.SalaryReport, error) {
	var results []models.SalaryReport
	monthName, err := normalizeMonth(month)
	if err != nil {
		return nil, err // Kembalikan error jika input bulan tidak valid
	}

	// Asumsikan kolom 's.month' di database Anda bertipe ANGKA (INT, SMALLINT, etc.)
	err = r.DB.Debug().Table("ak_salaries s").
		Select("s.user_uid, u.full_name, s.month, s.year, s.total_gaji, s.total_potongan").
		Joins("left join ak_users u on u.user_uid = s.user_uid").
		Where("s.month = ? AND s.year = ?", monthName, year). // Gunakan monthNumber di sini
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
