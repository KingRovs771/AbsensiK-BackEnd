package repository

import (
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
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

func (r *SalaryRepository) GetTotalIzin(userID string, month int, year int) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Ak_Izin{}).
		Where("user_uid = ? AND izin_tipe = 'izin'", userID).
		Where("EXTRACT(MONTH FROM start_date) = ?", month).
		Where("EXTRACT(YEAR FROM start_date) = ?", year).
		Where("status = 1").
		Count(&count).Error
	return count, err
}

func (r *SalaryRepository) GetTotalSakit(userID string, month int, year int) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Ak_Izin{}).
		Where("user_uid = ? AND izin_tipe = 'sakit'", userID).
		Where("EXTRACT(MONTH FROM start_date) = ?", month).
		Where("EXTRACT(YEAR FROM start_date) = ?", year).
		Where("status = 1").
		Count(&count).Error
	return count, err
}

func (r *SalaryRepository) GetTotalCuti(userID string, month int, year int) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Ak_Izin{}).
		Where("user_uid = ? AND izin_tipe = 'cuti'", userID).
		Where("EXTRACT(MONTH FROM start_date) = ?", month).
		Where("EXTRACT(YEAR FROM start_date) = ?", year).
		Where("status = 1").
		Count(&count).Error
	return count, err
}

// Mengambil total kehadiran dari Ak_kehadiran
func (r *SalaryRepository) GetTotalKehadiran(userID string, month int, year int) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Ak_Kehadiran{}).
		Where("user_uid = ?", userID).
		Where("EXTRACT(MONTH FROM tanggal) = ?", month).
		Where("EXTRACT(YEAR FROM tanggal) = ?", year).
		Count(&count).Error
	return count, err
}

// Mengambil total potongan dari Ak_potongans
func (r *SalaryRepository) GetTotalPotongan(userID string, month string, year int64) (int64, error) {
	var totalPotongan int64

	err := r.DB.Model(&models.Ak_Potongan{}).
		Select("COALESCE(SUM(tp.nilai_potongan), 0)").
		Joins("INNER JOIN ak_tipe_potongans tp ON ak_potongans.tipe_potongan = tp.tipe_potongan_id ").
		Where("ak_potongans.user_uid = ? AND ak_potongans.month = ? AND ak_potongans.year = ?", userID, month, year).
		Scan(&totalPotongan).Error

	return totalPotongan, err
}

// Mengambil daily rate berdasarkan User ID
func (r *SalaryRepository) GetDailyRate(userID string) (int64, error) {
	var dailyRate int64
	err := r.DB.Model(&models.Ak_Users{}).
		Where("user_uid = ?", userID).
		Select("dailyrate").Scan(&dailyRate).Error
	return dailyRate, err
}

func (r *SalaryRepository) GetSalariesByMonth(month string, year int) ([]models.Ak_Salary, error) {
	var salaries []models.Ak_Salary

	// Debugging: Cetak nilai bulan & tahun yang diterima
	fmt.Println("Repository Debug - Month:", month, "Year:", year)

	// Query database berdasarkan bulan & tahun yang dikirim frontend
	err := r.DB.Where("month = ? AND year = ?", month, year).Find(&salaries).Error
	return salaries, err
}

func (r *SalaryRepository) GetSalariesByMonthAndName(month string, year int) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	fmt.Println("Repository Debug - Month:", month, "Year:", year)

	// Query dengan JOIN untuk mengambil `full_name` berdasarkan `user_uid`
	err := r.DB.Table("ak_salaries").
		Select("ak_salaries.salary_id, ak_salaries.user_uid, ak_salaries.month, ak_salaries.year, ak_users.full_name").
		Joins("JOIN ak_users ON ak_salary.user_uid = ak_users.user_uid").
		Where("ak_salaries.month = ? AND ak_salaries.year = ?", month, year).
		Find(&results).Error

	if err != nil {
		fmt.Println("Error fetching salary data:", err)
		return nil, err
	}

	return results, nil
}
