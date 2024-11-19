package models

import "time"

type Ak_Izin struct {
	IzinId      uint64 		`gorm:"primaryKey" json:"izin_id"`
	UserId      string 		`gorm:"type:varchar(30)" json:"user_id"`
	IzinType    string 		`gorm:"type:varchar(20)" json:"izin_tipe"`
	StartDate   time.Time 	`gorm:"type:date" json:"start_date"`
	EndDate     time.Time 	`gorm:"type:date" json:"end_date"`
	Alasan 	    string 		`gorm:"type:text" json:"alasan"`
	Status 	    int 		`gorm:"type:int" json:"status"`
	ApproveBy   string 		`gorm:"type:varchar(30)" json:"approve_by"`
	ApproveDate time.Time 	`gorm:"type:date" json:"approve_date"`
}