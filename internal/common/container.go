package common

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	DB          *gorm.DB
	AuthHandler *handler.AuthHandler
}

func NewContainer(cfg *config.Config) (*Container, error) {
	db, err := database.NewConnection(cfg)
	if err != nil {
		return nil, err
	}
	
	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	return &Container{
		AuthHandler: authHandler,
	}, nil
}
