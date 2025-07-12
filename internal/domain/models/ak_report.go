package models

type SalaryReport struct {
	UserUID       string `json:"user_uid"`
	FullName      string `json:"full_name"`
	Month         string `json:"month"`
	Year          int64  `json:"year"`
	TotalGaji     int64  `json:"total_gaji"`
	TotalPotongan int64  `json:"total_potongan"`
}

type AttendanceReport struct {
	UserUID  string  `json:"user_uid"`
	FullName string  `json:"full_name"`
	Tanggal  string  `json:"tanggal"`
	TimeIn   *string `json:"time_in"`
	TimeOut  *string `json:"time_out"`
}

type LeaveReport struct {
	UserUID    string `json:"user_uid"`
	FullName   string `json:"full_name"`
	PermitType string `json:"permit_type"` // Izin, Sakit, atau Cuti
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`
	Status     string `json:"status"`
}
