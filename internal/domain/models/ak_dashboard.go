package models

type DashboardStats struct {
	TotalEmployees int64 `json:"total_employees"`
	OnLeaveToday   int64 `json:"on_leave_today"` // Karyawan Izin
	AbsentToday    int64 `json:"absent_today"`   // Karyawan Alpha
	PresentToday   int64 `json:"present_today"`  // Total Kehadiran
}
