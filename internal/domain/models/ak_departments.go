package models

type Ak_Departments struct {
	DepartmentsId   string `gorm:"primaryKey" json:"departments_id"`
	NameDepartments string `gorm:"type:varchar(30)" json:"name_departments"`
	Description     string `gorm:"type:text" json:"description"`
}
