package common

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handler"
	cateRepo "backend/internal/repository/category"
	otpRepo "backend/internal/repository/otp"
	userRepo "backend/internal/repository/user"
	auth "backend/internal/service/auth"
	category "backend/internal/service/category"

	"gorm.io/gorm"
)

type Container struct {
	DB              *gorm.DB
	AuthHandler     *handler.AuthHandler
	CategoryHandler *handler.CategoryHandler
	UserRepository  userRepo.UserRepository
}

func NewContainer(cfg *config.Config) (*Container, error) {
	db, err := database.NewConnection(cfg)
	if err != nil {
		return nil, err
	}

	userRepo := userRepo.NewUserRepository(db)
	otpRepo := otpRepo.NewOtpRepository(db)
	cateRepo := cateRepo.NewCategoryRepository(db)

	authService := auth.NewAuthService(userRepo, otpRepo)
	cateService := category.NewCategoryService(cateRepo)

	authHandler := handler.NewAuthHandler(authService, cfg.Server.JwtSecret)
	cateHandler := handler.NewCategoryHandler(cateService)

	return &Container{
		AuthHandler:     authHandler,
		CategoryHandler: cateHandler,
		UserRepository:  userRepo,
		DB:              db,
	}, nil
}
