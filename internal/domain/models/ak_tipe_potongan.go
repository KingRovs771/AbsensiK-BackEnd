package models

type Ak_TipePotongans struct {
	TipePotonganId int    `gorm:"primaryKey" json:"tipe_potongan_id"`
	NamePotongan   string `gorm:"type:varchar(70)" json:"name_potongan"`
	NilaiPotongan  int64  `gorm:"type:int" json:"nilai_potongan"`
}
