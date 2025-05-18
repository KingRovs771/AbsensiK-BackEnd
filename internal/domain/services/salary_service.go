package services

import (
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type SalaryService struct {
	Repo *repository.SalaryRepository
}

var monthMap = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

func NewSalaryService(repo *repository.SalaryRepository) *SalaryService {
	return &SalaryService{Repo: repo}
}

func (s *SalaryService) CheckAPI() string {
	return `message : Berhasil Menyambung`
}

// Helper Function untuk Error Handling
func handleError(err error, message string) map[string]interface{} {
	if err != nil {
		fmt.Println("Error:", err) // Debugging jika ada error
		return map[string]interface{}{
			"status":  "error",
			"message": message,
		}
	}
	return nil
}
func (s *SalaryService) GenerateSalary(userID, monthStr string, year int64) map[string]interface{} {
	month, exists := monthMap[monthStr]
	if !exists {
		return map[string]interface{}{
			"status":  "error",
			"message": "Format bulan tidak valid",
			"month":   monthStr,
			"userId":  userID,
		}
	}
	yearInt := int(year)

	// Ambil data kehadiran
	totalKehadiran, err := s.Repo.GetTotalKehadiran(userID, month, yearInt)
	if res := handleError(err, "Gagal mengambil data kehadiran"); res != nil {
		return res
	}

	// Ambil data izin, sakit, cuti
	totalIzin, err := s.Repo.GetTotalIzin(userID, month, yearInt)
	if res := handleError(err, "Gagal mengambil data izin"); res != nil {
		return res
	}

	totalSakit, err := s.Repo.GetTotalSakit(userID, month, yearInt)
	if res := handleError(err, "Gagal mengambil data sakit"); res != nil {
		return res
	}

	totalCuti, err := s.Repo.GetTotalCuti(userID, month, yearInt)
	if res := handleError(err, "Gagal mengambil data cuti"); res != nil {
		return res
	}

	// Ambil total potongan gaji
	totalPotongan, err := s.Repo.GetTotalPotongan(userID, monthStr, year)
	if res := handleError(err, "Gagal mengambil data potongan"); res != nil {
		return res
	}

	// Ambil daily rate
	dailyRate, err := s.Repo.GetDailyRate(userID)
	if res := handleError(err, "Gagal mengambil daily rate"); res != nil {
		return res
	}

	// Validasi daily rate tidak boleh nol untuk menghindari kesalahan perhitungan
	if dailyRate == 0 {
		return map[string]interface{}{
			"status":  "error",
			"message": "Daily rate tidak valid (0)",
			"userId":  userID,
		}
	}

	// Menghitung gaji akhir
	baseSalary := dailyRate * totalKehadiran
	totalSalary := baseSalary - totalPotongan

	// Simpan data gaji
	salary := models.Ak_Salary{
		UserUID:        userID,
		Month:          monthStr,
		Year:           year,
		TotalKehadiran: totalKehadiran,
		TotalIzin:      totalIzin,
		TotalSakit:     totalSakit,
		TotalCuti:      totalCuti,
		TotalPotongan:  totalPotongan,
		TotalGaji:      totalSalary,
	}

	err = s.Repo.SaveSalary(&salary)
	if res := handleError(err, "Gagal menyimpan data gaji"); res != nil {
		return res
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Gaji berhasil dihitung",
		"data":    salary,
	}
}

func (s *SalaryService) GetSalariesByMonth(month string, year int) map[string]interface{} {
	salaries, err := s.Repo.GetSalariesByMonth(month, year)
	if err != nil {
		fmt.Println("Service Debug - Error Fetching Salaries:", err)
		return map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve salary data",
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Salary data retrieved successfully",
		"Data":    salaries,
	}
}
func (s *SalaryService) GetSalariesByMonthAndName(month string, year int) map[string]interface{} {
	salaries, err := s.Repo.GetSalariesByMonthAndName(month, year)
	if err != nil {
		fmt.Println("Service Debug - Error Fetching Salaries:", err)
		return map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve salary data",
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Salary data retrieved successfully",
		"Data":    salaries,
	}
}
