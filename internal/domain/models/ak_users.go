package models

type AK_USERS struct {
	UserId        int64  `gorm:"primaryKey" json:"user_id"`
	UserUID       string `gorm:"type:varchar(90)" json:"user_uid"`
	Username      string `gorm:"type:varchar(80)" json:"username"`
	Password      string `gorm:"type:varchar(100)" json:"password"`
	Email         string `gorm:"type:varchar(80)" json:"email"`
	FullName      string `gorm:"type:varchar(80)" json:"full_name"`
	Gender        string `gorm:"type:varchar(20)" json:"gender"`
	Phone         string `gorm:"type:varchar(15)" json:"phone"`
	Address       string `gorm:"type:text" json:"address"`
	Dailyrate     int64  `gorm:"type:int" json:"dailyrate"`
	DepartementId string `gorm:"type:string" json:"departement_id"`
	RoleId        string `gorm:"type:varchar(15)" json:"role_id"`
}