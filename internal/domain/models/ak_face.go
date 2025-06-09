package models

type Ak_Face struct {
	FacesID  int64  `gorm:"primaryKey" json:"faces_id"`
	UserUID  string `gorm:"type:varchar(30)" json:"user_uid"`
	FaceData []byte `gorm:"type:longblob;not null" json:"face_data"`

	User Ak_Users `gorm:"foreignKey:UserUID;references:UserUID" json:"user"`
}
