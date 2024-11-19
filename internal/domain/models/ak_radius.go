package models

type Ak_Radius struct {
	RadiusId     int     `gorm:"primaryKey" json:"radius_id"`
	NameLocation string  `gorm:"type:varchar(70)" json:"name_location"`
	Latitude     float64 `gorm:"type:decimal(9,6)" json:"latitude"`
	Longitude    float64 `gorm:"type:decimal(9,6)" json:"longitude"`
	Radius       int     `gorm:"type:int" json:"radius"`
}