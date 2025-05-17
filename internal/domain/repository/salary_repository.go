package repository

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
	"time"
)

type SalaryRepository struct {
	DB *gorm.DB
}

func NewSalaryRepository(db *gorm.DB) *SalaryRepository {
	return &SalaryRepository{DB: db}
}

// Simpan data gaji ke database
func (r *SalaryRepository) SaveSalary(salary *models.Ak_Salary) error {
	return r.DB.Create(salary).Error
}

// Mengambil jumlah izin berdasarkan Ak_izins
func (r *SalaryRepository) GetTotalIzin(userID, month string, year int64) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Ak_Izin{}).
		Where("user_id = ? AND izin_tipe = 'izin' AND month = ? AND year = ?", userID, month, year).
		Count(&count).Error
	return count, err
}

func (r *SalaryRepository) GetTotalSakit(userID, month string, year int64) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Ak_Izin{}).
		Where("user_id = ? AND izin_tipe = 'sakit' AND month = ? AND year = ?", userID, month, year).
		Count(&count).Error
	return count, err
}

// Mengambil jumlah cuti berdasarkan Ak_izins
func (r *SalaryRepository) GetTotalCuti(userID, month string, year int64) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Ak_Izin{}).
		Where("user_id = ? AND izin_tipe = 'cuti' AND month = ? AND year = ?", userID, month, year).
		Count(&count).Error
	return count, err
}

// Mengambil total kehadiran dari Ak_kehadiran
func (r *SalaryRepository) GetTotalKehadiran(userID, month string, year int64) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Ak_Kehadiran{}).
		Where("user_id = ? AND month = ? AND year = ?", userID, month, year).
		Count(&count).Error
	return count, err
}

// Mengambil total potongan dari Ak_potongans
func (r *SalaryRepository) GetTotalPotongan(userID, month string, year int64) (int64, error) {
	var totalPotongan int64
	err := r.DB.Model(&models.Ak_Potongan{}).
		Where("user_id = ? AND month = ? AND year = ?", userID, month, year).
		Select("SUM(amount)").Scan(&totalPotongan).Error
	return totalPotongan, err
}

// Mengambil daily rate berdasarkan User ID
func (r *SalaryRepository) GetDailyRate(userID string) (int64, error) {
	var dailyRate int64
	err := r.DB.Model(&models.Ak_Users{}).
		Where("user_id = ?", userID).
		Select("daily_rate").Scan(&dailyRate).Error
	return dailyRate, err
}

func (r *SalaryRepository) GetSalariesByCurrentMonth() ([]models.Ak_Salary, error) {
	var salaries []models.Ak_Salary

	// Mendapatkan bulan dan tahun saat ini
	currentTime := time.Now()
	currentMonth := currentTime.Format("January") // Contoh: "May"
	currentYear := int64(currentTime.Year())      // Contoh: 2025

	// Query database berdasarkan bulan dan tahun saat ini
	err := r.DB.Where("month = ? AND year = ?", currentMonth, currentYear).Find(&salaries).Error
	return salaries, err
}
