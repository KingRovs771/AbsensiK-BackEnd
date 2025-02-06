package models

type Ak_Radius struct {
	RadiusId     int    `gorm:"primaryKey" json:"radius_id"`
	NameLocation string `gorm:"type:varchar(70)" json:"name_location"`
	Latitude     string `gorm:"type:varchar(200)" json:"latitude"`
	Longitude    string `gorm:"type:varchar(200)" json:"longitude"`
	Radius       int    `gorm:"type:int" json:"radius"`
}
