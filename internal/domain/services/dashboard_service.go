package services

import (
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
	"log"
	"time"
)

type DashboardService interface {
	GetDashboardStats() (*models.DashboardStats, error)
}

type dashboardService struct {
	repo repository.DashboardRepository
}

func NewDashboardService(repo repository.DashboardRepository) DashboardService {
	return &dashboardService{repo: repo}
}

func (s *dashboardService) GetDashboardStats() (*models.DashboardStats, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	todayDate := now.Format("2006-01-02")
	todayDay := now.Weekday().String()

	totalEmployees, err := s.repo.CountTotalUsers()
	if err != nil {
		log.Printf("Error counting total users: %v", err)
		return nil, err
	}

	onLeave, err := s.repo.CountOnLeaveToday(todayDate)
	if err != nil {
		log.Printf("Error counting users on leave: %v", err)
		return nil, err
	}

	present, err := s.repo.CountPresentToday(todayDate)
	if err != nil {
		log.Printf("Error counting present users: %v", err)
		return nil, err
	}

	absent, err := s.repo.CountAbsentToday(todayDate, todayDay)
	if err != nil {
		log.Printf("Error counting absent users: %v", err)
		return nil, err
	}

	stats := &models.DashboardStats{
		TotalEmployees: totalEmployees,
		OnLeaveToday:   onLeave,
		AbsentToday:    absent,
		PresentToday:   present,
	}

	return stats, nil
}
