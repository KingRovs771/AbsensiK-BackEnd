package models

type Ak_Users struct {
	UserId        int64          `gorm:"primaryKey" json:"user_id"`
	UserUID       string         `gorm:"type:varchar(90)" json:"user_uid"`
	Username      string         `gorm:"type:varchar(80)" json:"username"`
	Password      string         `gorm:"type:varchar(100)" json:"password"`
	Email         string         `gorm:"type:varchar(80)" json:"email"`
	FullName      string         `gorm:"type:varchar(80)" json:"full_name"`
	Gender        string         `gorm:"type:varchar(20)" json:"gender"`
	Phone         string         `gorm:"type:varchar(15)" json:"phone"`
	Address       string         `gorm:"type:text" json:"address"`
	Dailyrate     int64          `gorm:"type:int" json:"dailyrate"`
	DepartmentsId string         `gorm:"type:varchar(15)" json:"departments_id"`
	Department    Ak_Departments `gorm:"foreignKey:DepartmentsId;references:DepartmentsId" json:"department"`
	RoleId        string         `gorm:"type:varchar(15)" json:"role_id"`
	Role          Ak_Roles       `gorm:"foreignKey:RoleId;references:RoleId" json:"role"`
}
