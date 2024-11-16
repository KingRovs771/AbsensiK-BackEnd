package database

import (
	"database/sql"

	"github.com/KingRovs771/AbsensiK-BackEnd/config"
)

func NewDatabase(cfg *config.Config) *sql.DB{
	connection := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" +cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName

	db, err :=sql.Open("Mysql", connection)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}