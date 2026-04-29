package api

import (
	"fmt"
	"inno-accounting/internal/adapters/jwt"
	"inno-accounting/internal/adapters/postgres"
	"inno-accounting/internal/adapters/repositories"
	"inno-accounting/internal/controllers"
	"inno-accounting/internal/controllers/handlers"
	"inno-accounting/internal/use-cases/auth"
	"inno-accounting/internal/use-cases/item/document"
	"inno-accounting/internal/use-cases/item/software"
	"inno-accounting/internal/use-cases/item/tech"
	storages "inno-accounting/internal/use-cases/storage"
	"inno-accounting/internal/use-cases/user"
	jwt_config "inno-accounting/pkg/jwt"
	"inno-accounting/pkg/logger"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type API struct {
	config *Config
}

func New(config *Config) *API {
	return &API{
		config: config,
	}
}

func (a *API) Start() error {

	// LOGGER INIT
	if err := logger.InitLogger(a.config.LoggerLevel); err!= nil {
		log.Print("Cannot read logger level")
		return err
	}
	logger.Info("Logger was configure successfully")

	// POSTGRES INIT
	storage, err := postgres.Init(a.config.PostgresURI)
	if err != nil {
		return fmt.Errorf("failed to connect db, %w", err)
	}
	logger.Info("Storage (postgres) was configure successfully")

	// JWT MANAGER INIT
	jwtCfg := jwt_config.New(
		[]byte(a.config.JWT.AccessSecret),
		[]byte(a.config.JWT.RefreshSecret),
		time.Duration(a.config.JWT.AccessTTL)*time.Minute,
		time.Duration(a.config.JWT.RefreshTTL)*time.Hour,
	)

	jwtManager := jwt.NewJWTTokenManager(jwtCfg)

	// BUSINESS LOGIC INIT
	userRepo := repositories.NewUserRepository(storage.GetDB())
	userService := user.New(userRepo)

	authRepo := repositories.NewAuthRepository(storage.GetDB())
	authService := auth.New(authRepo, userService, jwtManager)

	storageRepo := repositories.NewStorageRepository(storage.GetDB())
	storageService := storages.New(storageRepo)

	techRepo := repositories.NewTechRepository(storage.GetDB())
	techService := tech.New(techRepo, storageService, userService) 

	docsRepo := repositories.NewDocumentRepository(storage.GetDB())
	docsService := document.New(docsRepo, storageService, userService)

	softwareRepo := repositories.NewSoftwareRepository(storage.GetDB())
	softwareService := software.New(softwareRepo, userService)

	// ROUTER INIT
	handlers := handlers.New(userService, authService, storageService, techService, docsService, softwareService)
	router := controllers.InitRouter(handlers, jwtManager)
	
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(router)
	logger.Info("Router was configure successfully")

	// SERVER RUNNING
	logger.Info("API STARTED AT PORT", a.config.BindAddr)
	return http.ListenAndServe(a.config.BindAddr, handler)
}