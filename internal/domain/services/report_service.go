package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
)

type ReportService interface {
	GetSalaryReport(month, year string) ([]models.SalaryReport, error)
	GetAttendanceReport(month, year string) ([]models.AttendanceReport, error)
	GetLeaveReport(month, year string) ([]models.LeaveReport, error)
	GetAllUsers() ([]models.Ak_Users, error)
}

type reportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) GetSalaryReport(month, year string) ([]models.SalaryReport, error) {
	return s.repo.GetSalaryReport(month, year)
}
func (s *reportService) GetAttendanceReport(month, year string) ([]models.AttendanceReport, error) {
	return s.repo.GetAttendanceReport(month, year)
}
func (s *reportService) GetLeaveReport(month, year string) ([]models.LeaveReport, error) {
	return s.repo.GetLeaveReport(month, year)
}
func (s *reportService) GetAllUsers() ([]models.Ak_Users, error) {
	return s.repo.GetAllUsers()
}
