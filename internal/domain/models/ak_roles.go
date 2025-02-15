package models

type Ak_Roles struct {
	RoleId      string `gorm:"primaryKey" json:"role_id"`
	NameRole    string `gorm:"type:varchar(30)" json:"name_role"`
	Description string `gorm:"type:text" json:"description"`
}
