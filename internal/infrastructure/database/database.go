package database

import (
	"fmt"
	"log"

	"github.com/KingRovs771/AbsensiK-BackEnd/config"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(config *config.Config) *gorm.DB{
	err := godotenv.Load()

	if err !=nil{
		log.Fatal("Error Loading .env file")
		log.Println("ERROR Loading ENV FILE")
	}

	dsn := config.DBUser + ":" + config.DBPassword + 
	"@tcp(" + config.DBHost + ":" + config.DBPort + 
	")/" + config.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err :=db.DB()
	if err != nil{
		log.Fatal(err)
	}
	sqlDB.SetMaxOpenConns(10)

	fmt.Println("Database Connected Succesfull")
	return db
}