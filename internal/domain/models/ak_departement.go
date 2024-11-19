package models

type Ak_Departement struct {
	DepartementId   string `gorm:"primaryKey" json:"departement_id"`
	NameDepartement string `gorm:"type:varchar(30) json:"name_departement"`
	Description     string `gorm:"type:text" json:"description"`
}