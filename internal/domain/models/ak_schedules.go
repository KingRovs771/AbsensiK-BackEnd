package models

import "time"

type Ak_Schedules struct {
	SchedulesId int64     `gorm:"primaryKey" json:"schedules_id"`
	UserId      string    `gorm:"type:varchar(30)" json:"user_id"`
	StartTime   time.Time `gorm:"type:time" json:"start_time"`
	EndTime     time.Time `gorm:"type:time" json:"end_time"`
	Day         string    `gorm:"type:varchar(30)" json:"day"`
	IsActive    int       `gorm:"type:int" json:"is_active"`
}
