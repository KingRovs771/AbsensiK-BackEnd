package models

type Ak_Salary struct {
	SalaryId       int64  `gorm:"primaryKey" json:"salary_id"`
	UserUID        string `gorm:"type:varchar(30)" json:"user_uid"`
	Month          string `gorm:"type:varchar(15)" json:"month"`
	Year           int64  `gorm:"type:int" json:"year"`
	TotalKehadiran int64  `gorm:"type:int" json:"total_kehadiran"`
	TotalIzin      int64  `gorm:"type:int" json:"total_izin"`
	TotalSakit     int64  `gorm:"type:int" json:"total_sakit"`
	TotalCuti      int64  `gorm:"type:int" json:"total_cuti"`
	TotalPotongan  int64  `gorm:"type:int" json:"total_potongan"`
	TotalGaji      int64  `gorm:"type:int" json:"total_gaji"`

	Users      Ak_Users       `gorm:"foreignKey:UserUID;references:UserUID" json:"users"`
	Role       Ak_Roles       `gorm:"foreignKey:RoleId;references:RoleId" json:"role"`
	Department Ak_Departments `gorm:"foreignKey:DepartmentsId;references:DepartmentsId" json:"department"`
}
