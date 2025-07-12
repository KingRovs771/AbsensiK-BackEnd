package models

import "time"

type Ak_Kehadiran struct {
	KehadiranId int64     `gorm:"type:primaryKey" json:"kehadiran_id"`
	UserUID     string    `gorm:"type:varchar(30)" json:"user_uid"`
	ScheduleId  string    `gorm:"type:varchar(30)" json:"schedule_id"`
	Tanggal     string    `gorm:"type:date" json:"tanggal"`
	TimeIn      *string   `gorm:"type:time" json:"time_in"`
	TimeOut     *string   `gorm:"type:time" json:"time_out"`
	Photo       []byte    `gorm:"type:varchar(100)" json:"photo"`
	Latitude    float64   `gorm:"type:decimal(9,6)" json:"latitude"`
	Longitude   float64   `gorm:"type:decimal(9,6)" json:"Longitude"`
	Radius      float64   `gorm:"type:int" json:"radius"`
	CreatedAt   time.Time `gorm:"type:date" json:"created_at"`
}
