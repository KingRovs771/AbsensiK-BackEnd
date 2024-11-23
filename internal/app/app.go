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

	configs := config.LoadConfig()
	db := database.NewDatabase(configs)
	// # Repo
	userRepo := repository.NewUserRepository(db)
	authRepo := repository.NewAuthRepostory(db)
	DepartementsRepo := repository.NewDepartementsRepository(db)

	// # Service
	userService := services.NewUserService(userRepo, "secretyaa")
	authService := services.NewAuthService(authRepo, "secretyaa")
	departementService := services.NewDepartementService(DepartementsRepo, "secretyaa")
	// # User Handler
	userHandler := deliveryhttp.NewUserHandler(userService)
	authHandler := deliveryhttp.NewAuthHandler(authService)
	departementHandler := deliveryhttp.NewDepartementHandler(departementService)
	// # Auth Handlers
	router := router.NewRouter(userHandler, authHandler, departementHandler)

	return &App{Router: router}
}

func (a *App) SetupRouter() http.Handler {
	return a.Router
}
