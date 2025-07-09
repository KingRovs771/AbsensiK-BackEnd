package models

type TodayAttendanceInfo struct {
	TimeIn  *string `json:"time_in"`
	TimeOut *string `json:"time_out"`
}
