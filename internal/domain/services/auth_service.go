package services

import (
	"log"
	"time"

	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/models"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
	SecretKey string
}

func NewAuthService(authRepo *repository.AuthRepository, secretKey string) *AuthService{
	return &AuthService{
		AuthRepository: authRepo,
		SecretKey: secretKey,
	}
}

func (s *AuthService)GenerateToken(user *models.AK_USERS)(string, error){
	claims := jwt.MapClaims{
		"user_id" : user.UserId,
		"username" : user.Username,
		"exp" : time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		log.Println("Error Generating Token :", err)
		return "", err
	}
	log.Println("Token Sukses Di Generate for User :", user.Username)
	return tokenString, nil
}

func (s *AuthService) Authenticate(username, password string)(string, error){
	user, err := s.AuthRepository.GetUserByUsername(username)
	if err != nil {
		log.Println("Error Fetching User: ", err)
		return "", err
	}

	if user == nil{
		log.Println("User Not Found :", username)
		return "", gorm.ErrRecordNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		return "", err
	}


	token, err :=s.GenerateToken(user)
	if err != nil {
		log.Println("Error Generating Token : ", err)
		return "", err
	}
	return token, nil
}