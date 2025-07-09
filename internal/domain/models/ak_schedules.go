package models

type Ak_Schedules struct {
	ScheduleId int64  `gorm:"primaryKey" json:"schedule_id"`
	UserUID    string `gorm:"type:varchar(30)" json:"user_uid"`
	StartTime  string `gorm:"type:time" json:"start_time"`
	EndTime    string `gorm:"type:time" json:"end_time"`
	Day        string `gorm:"type:varchar(30)" json:"day"`
	IsActive   int    `gorm:"type:int" json:"is_active"`

	Users Ak_Users `gorm:"foreignKey:UserUID;references:UserUID" json:"user"`
}
