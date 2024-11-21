package repository

import (
	"log"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepostory(db *gorm.DB) *AuthRepository{
	return &AuthRepository{DB : db}
}

func (r *AuthRepository) GetUserByUsername(username string)(*models.AK_USERS, error){
	var Users models.AK_USERS
	err:=r.DB.Where("username = ?", username).First(&Users).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			log.Println("User Not Found :", username)
			return nil, nil
		}
		log.Println("Error Fetching User :", err)
		return nil, err
	}
	log.Println("User Fetched success :", Users.Username)
	return &Users, nil
}