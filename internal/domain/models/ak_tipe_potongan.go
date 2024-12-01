package models

type Ak_TipePotongan struct {
	TipePotonganId int64 `gorm:"type:primaryKey" json:"tipe_potongan_id"`
	NamaPotongan string `gorm:"type:varchar(70)" json:"nama_potongan"`
	NilaiPotongan int64 `gorm:"type:int" json:"nilai_potongan"`
}