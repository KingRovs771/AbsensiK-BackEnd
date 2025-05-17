package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type SalaryService struct {
	Repo *repository.SalaryRepository
}

func NewSalaryService(repo *repository.SalaryRepository) *SalaryService {
	return &SalaryService{Repo: repo}
}

func (s *SalaryService) CheckAPI() string {
	return `message : Berhasil Menyambung`
}

func (s *SalaryService) GenerateSalary(userID, month string, year int64) map[string]interface{} {
	totalKehadiran, err := s.Repo.GetTotalKehadiran(userID, month, year)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Gagal mengambil data kehadiran",
		}
	}

	totalIzin, err := s.Repo.GetTotalIzin(userID, month, year)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Gagal mengambil data izin",
		}
	}

	totalSakit, err := s.Repo.GetTotalSakit(userID, month, year)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Gagal mengambil data sakit",
		}
	}

	totalCuti, err := s.Repo.GetTotalCuti(userID, month, year)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Gagal mengambil data cuti",
		}
	}

	totalPotongan, err := s.Repo.GetTotalPotongan(userID, month, year)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Gagal mengambil data potongan",
		}
	}

	dailyRate, err := s.Repo.GetDailyRate(userID)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Gagal mengambil daily rate",
		}
	}

	// Menghitung gaji akhir
	baseSalary := dailyRate * totalKehadiran
	totalSalary := baseSalary - totalPotongan

	salary := models.Ak_Salary{
		UserUID:        userID,
		Month:          month,
		Year:           year,
		TotalKehadiran: totalKehadiran,
		TotalIzin:      totalIzin,
		TotalSakit:     totalSakit,
		TotalCuti:      totalCuti,
		TotalPotongan:  totalPotongan,
		TotalGaji:      totalSalary,
	}

	err = s.Repo.SaveSalary(&salary)
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Gagal menyimpan data gaji",
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Gaji berhasil dihitung",
		"data":    salary,
	}
}

func (s *SalaryService) GetSalariesByCurrentMonth() map[string]interface{} {
	salaries, err := s.Repo.GetSalariesByCurrentMonth()
	if err != nil {
		return map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve salary data",
		}
	}

	return map[string]interface{}{
		"status":  "success",
		"message": "Salary data retrieved successfully",
		"data":    salaries,
	}
}
