package services

import (
	"fmt"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
	"gorm.io/gorm"
	"io"
	"math"
	"mime/multipart"
	"strconv"
	"time"
)

type attendanceService struct {
	repo *repository.AttendanceRepository
}

func NewAttendanceService(repo *repository.AttendanceRepository) *AttendanceService {
	return &attendanceService{repo: repo}
}

func (s *attendanceService) GetAttendancePageData(userUID string) (*models.AttendancePageData, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	todayDate := now.Format("2006-01-02")
	todayDay := now.Format("Monday")

	schedule, err := s.repo.GetTodaySchedule(userUID, todayDay)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &models.AttendancePageData{
				Schedule:    models.ScheduleInfo{StartTime: "LIBUR", EndTime: "LIBUR"},
				CanClockIn:  false,
				CanClockOut: false,
			}, nil
		}
		return nil, err
	}

	officeLocation, err := s.repo.GetOfficeLocation()
	if err != nil {
		return nil, err
	}

	todayAttendance, err := s.repo.GetTodayAttendance(userUID, todayDate)

	data := &models.AttendancePageData{
		Schedule: models.ScheduleInfo{
			StartTime: schedule.StartTime,
			EndTime:   schedule.EndTime,
		},
		OfficeLocation: models.OfficeLocationInfo{
			Latitude:  officeLocation.Latitude,
			Longitude: officeLocation.Longitude,
			Radius:    officeLocation.Radius,
		},
		Attendance:         models.TodayAttendanceInfo{},
		LateDuration:       "0 menit",
		EarlyLeaveDuration: "0 menit",
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if todayAttendance != nil {
		data.Attendance.TimeIn = todayAttendance.TimeIn
		data.Attendance.TimeOut = todayAttendance.TimeOut

		if todayAttendance.TimeIn != nil {
			lateDuration, _ := calculateDuration(*todayAttendance.TimeIn, schedule.StartTime, true)
			data.LateDuration = lateDuration
		}
		if todayAttendance.TimeOut != nil {
			earlyDuration, _ := calculateDuration(*todayAttendance.TimeOut, schedule.EndTime, false)
			data.EarlyLeaveDuration = earlyDuration
		}
	}

	data.CanClockIn = (todayAttendance == nil || todayAttendance.TimeIn == nil)
	data.CanClockOut = (todayAttendance != nil && todayAttendance.TimeIn != nil && todayAttendance.TimeOut == nil)

	return data, nil
}

func (s *attendanceService) PerformClockIn(userUID string, photoFile multipart.File, latStr, lonStr string) (string, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	todayDate := now.Format("2006-01-02")
	currentTime := now.Format("15:04:05")
	todayDay := now.Format("Monday")

	_, err := s.repo.GetTodayAttendance(userUID, todayDate)
	if err == nil {
		return "", fmt.Errorf("Anda sudah melakukan absen masuk hari ini")
	}
	if err != gorm.ErrRecordNotFound {
		return "", err
	}

	officeLocation, err := s.repo.GetOfficeLocation()
	if err != nil {
		return "", fmt.Errorf("gagal mendapatkan lokasi kantor")
	}
	lat, _ := strconv.ParseFloat(latStr, 64)
	lon, _ := strconv.ParseFloat(lonStr, 64)
	distance := calculateDistance(lat, lon, officeLocation.Latitude, officeLocation.Longitude)

	// PERBAIKAN DI SINI: Casting officeLocation.Radius ke float64
	if distance > float64(officeLocation.Radius) {
		return "", fmt.Errorf("Anda berada di luar radius kantor (Jarak: %.f meter)", distance)
	}

	photoBytes, err := io.ReadAll(photoFile)
	if err != nil {
		return "", fmt.Errorf("gagal membaca file foto")
	}

	schedule, err := s.repo.GetTodaySchedule(userUID, todayDay)
	if err != nil {
		return "", fmt.Errorf("tidak ada jadwal kerja aktif untuk hari ini")
	}

	newAttendance := &models.Ak_Kehadiran{
		UserUID:    userUID,
		ScheduleId: strconv.FormatInt(schedule.ScheduleId, 10),
		Tanggal:    todayDate,
		TimeIn:     &currentTime,
		Photo:      photoBytes,
		Latitude:   lat,
		Longitude:  lon,
		Radius:     int64(officeLocation.Radius),
	}

	if err := s.repo.CreateClockIn(newAttendance); err != nil {
		return "", err
	}

	return "Absen Masuk berhasil direkam", nil
}

func (s *attendanceService) PerformClockOut(userUID string, latStr, lonStr string) (string, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	todayDate := now.Format("2006-01-02")
	currentTime := now.Format("15:04:05")

	attendance, err := s.repo.GetTodayAttendance(userUID, todayDate)
	if err == gorm.ErrRecordNotFound {
		return "", fmt.Errorf("Anda belum melakukan absen masuk hari ini")
	}
	if err != nil {
		return "", err
	}
	if attendance.TimeOut != nil {
		return "", fmt.Errorf("Anda sudah melakukan absen pulang hari ini")
	}

	officeLocation, err := s.repo.GetOfficeLocation()
	if err != nil {
		return "", fmt.Errorf("gagal mendapatkan lokasi kantor")
	}
	lat, _ := strconv.ParseFloat(latStr, 64)
	lon, _ := strconv.ParseFloat(lonStr, 64)
	distance := calculateDistance(lat, lon, officeLocation.Latitude, officeLocation.Longitude)

	// PERBAIKAN DI SINI: Casting officeLocation.Radius ke float64
	if distance > float64(officeLocation.Radius) {
		return "", fmt.Errorf("Anda berada di luar radius kantor (Jarak: %.f meter)", distance)
	}

	if err := s.repo.UpdateClockOut(userUID, todayDate, currentTime); err != nil {
		return "", err
	}

	return "Absen Pulang berhasil direkam", nil
}

func (s *attendanceService) GetAllAttendancesForToday() ([]models.Ak_Kehadiran, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	todayDate := time.Now().In(loc).Format("2006-01-02")

	return s.repo.GetAllAttendancesWithUser(todayDate)
}

// --- Fungsi Helper ---

func calculateDuration(timeStr, scheduleTimeStr string, isLateCheck bool) (string, error) {
	layout := "15:04:05"
	actualTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return "", err
	}
	scheduleTime, err := time.Parse(layout, scheduleTimeStr)
	if err != nil {
		return "", err
	}

	var diff time.Duration
	if isLateCheck {
		if actualTime.After(scheduleTime) {
			diff = actualTime.Sub(scheduleTime)
		}
	} else {
		if actualTime.Before(scheduleTime) {
			diff = scheduleTime.Sub(actualTime)
		}
	}

	if diff.Minutes() < 1 {
		return "0 menit", nil
	}
	return fmt.Sprintf("%.0f menit", diff.Minutes()), nil
}

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000
	radLat1 := lat1 * math.Pi / 180
	radLat2 := lat2 * math.Pi / 180
	deltaLat := (lat2 - lat1) * math.Pi / 180
	deltaLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(radLat1)*math.Cos(radLat2)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
