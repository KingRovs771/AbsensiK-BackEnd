package models

import "time"

type Ak_Kehadiran struct {
	KehadiranId int64     `gorm:"type:primaryKey" json:"kehadiran_id"`
	UserId      string    `gorm:"type:varchar(30)" json:"user_id"`
	ScheduleId  string    `gorm:"type:varchar(30)" json:"schedule_id"`
	Tanggal     time.Time `gorm:"type:date" json:"tanggal"`
	TimeIn      time.Time `gorm:"type:time" json:"time_in"`
	TimeOut     time.Time `gorm:"type:time" json:"time_out"`
	Photo       string    `gorm:"type:varchar(100)" json:"photo"`
	Latitude    float64   `gorm:"type:decimal(9,6)" json:"latitude"`
	Longitude   float64   `gorm:"type:decimal(9,6)" json:"Longitude"`
	Radius      int64     `gorm:"type:int" json:"radius"`
}
