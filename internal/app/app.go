package app

import (
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/config"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/infrastructure/database"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/infrastructure/router"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/infrastructure/services"
)

type App struct {
	Router http.Handler
}

func NewApp() *App {

	config := config.LoadConfig()
	db := database.NewDatabase(config)

	userService := services.NewUserService(db)
	userHandler := http.NewUserHandler(userService)

	router := router.NewRouter(userHandler)

	return &App{Router: router}
}

func (a *App) SetupRouter() http.Handler{
	return a.Router
}