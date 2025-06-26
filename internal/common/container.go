package common

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handler"
	otpRepo "backend/internal/repository/otp"
	userRepo "backend/internal/repository/user"
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

	userRepo := userRepo.NewUserRepository(db)
	otpRepo := otpRepo.NewOtpRepository(db)
	authService := service.NewAuthService(userRepo, otpRepo)

	authHandler := handler.NewAuthHandler(authService)

	return &Container{
		AuthHandler: authHandler,
	}, nil
}
