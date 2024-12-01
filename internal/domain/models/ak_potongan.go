package models

type Ak_Potongan struct {
	PotonganId   int64  `gorm:"type:primaryKey" json:"potongan_id"`
	UserId       string `gorm:"type:varchar(30)" json:"user_id"`
	Month        string `gorm:"type:varchar(15)" json:"month"`
	Year         int64  `gorm:"type:int" json:"year"`
	TipePotongan string `gorm:"type:varchar(20)" json:"tipe_potongan"`
}
