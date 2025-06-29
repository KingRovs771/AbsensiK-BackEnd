package models

import "time"

type Ak_Izin struct {
	IzinId      int64     `gorm:"primaryKey" json:"izin_id"`
	UserUID     string    `gorm:"type:varchar(30)" json:"user_uid"`
	IzinType    string    `gorm:"type:varchar(20)" json:"izin_tipe"`
	StartDate   string    `gorm:"type:date" json:"start_date"`
	EndDate     string    `gorm:"type:date" json:"end_date"`
	Alasan      string    `gorm:"type:text" json:"alasan"`
	Status      int       `gorm:"type:int" json:"status"`
	Foto        []byte    `gorm:"type:longblob;" json:"foto"`
	ApproveBy   string    `gorm:"type:varchar(30)" json:"approve_by"`
	ApproveDate time.Time `gorm:"type:date" json:"approve_date"`

	Users Ak_Users `gorm:"foreignKey:UserUID;references:UserUID" json:"users"`
}
