package models

import "time"

type AK_USERS struct {
	UserId         string   `gorm:"primaryKey" json:"user_id"`
	UserUID        string  `gorm:"type:varchar(90)" json:"user_uid"`
	Username       string  `gorm:"type:varchar(80)" json:"username"`
	Password       string  `gorm:"type:varchar(100)" json:"password"`
	Email          string  `gorm:"type:varchar(80)" json:"email"`
	FullName       string  `gorm:"type:varchar(80)" json:"full_name"`
	Gender         string  `gorm:"type:varchar(20)" json:"gender"`
	Phone          string  `gorm:"type:varchar(15)" json:"phone"`
	Address        string  `gorm:"type:text" json:"address"`
	Dailyrate      float64 `gorm:"type:decimal" json:"dailyrate"`
	DepartementId  string  `gorm:"type:string" json:"departement_id"`
	RoleId         string  `gorm:"type:varchar(15)" json:"role_id"`
	Created_at     time.Time `gorm:"autoCreateTime" json:"created_at"`
	Update_at      time.Time `gorm:"autoCreateTime" json:"update_at"`
}