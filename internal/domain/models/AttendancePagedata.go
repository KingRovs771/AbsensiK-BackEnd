package models

type AttendancePageData struct {
	Schedule           ScheduleInfo        `json:"schedule"`
	OfficeLocation     OfficeLocationInfo  `json:"office_location"`
	Attendance         TodayAttendanceInfo `json:"attendance"`
	LateDuration       string              `json:"late_duration"`
	EarlyLeaveDuration string              `json:"early_leave_duration"`
	CanClockIn         bool                `json:"can_clock_in"`
	CanClockOut        bool                `json:"can_clock_out"`
}
