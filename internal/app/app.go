package app

import (
	"net/http"

	"github.com/KingRovs771/AbsensiK-BackEnd/config"
	deliveryhttp "github.com/KingRovs771/AbsensiK-BackEnd/internal/delivery/http"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/repository"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/domain/services"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/infrastructure/database"
	"github.com/KingRovs771/AbsensiK-BackEnd/internal/infrastructure/router"
)

type App struct {
	Router http.Handler
}

func NewApp() *App {

	configs:= config.LoadConfig()
	db:= database.NewDatabase(configs)
	// # Repo
	userRepo:= repository.NewUserRepository(db)
	
	// # Service
	userService:= services.NewUserService(userRepo)

	// # User Handler
	userHandler := deliveryhttp.NewUserHandler(userService)

	// # Auth Handlers
 	router:= router.NewRouter(userHandler)

	return &App{Router: router}
}

func (a *App) SetupRouter() http.Handler{
	return a.Router
}