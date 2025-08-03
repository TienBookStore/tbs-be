package common

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/handler"
	bookRepo "backend/internal/repository/book"
	cateRepo "backend/internal/repository/category"
	otpRepo "backend/internal/repository/otp"
	userRepo "backend/internal/repository/user"
	auth "backend/internal/service/auth"
	book "backend/internal/service/book"
	category "backend/internal/service/category"

	"gorm.io/gorm"
)

type Container struct {
	DB              *gorm.DB
	AuthHandler     *handler.AuthHandler
	CategoryHandler *handler.CategoryHandler
	BookHandler     *handler.BookHandler
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
	bookRepo := bookRepo.NewBookRepository(db)

	authService := auth.NewAuthService(userRepo, otpRepo)
	cateService := category.NewCategoryService(cateRepo)
	bookService := book.NewBookService(bookRepo, cateRepo)

	authHandler := handler.NewAuthHandler(authService, cfg.Server.JwtSecret)
	cateHandler := handler.NewCategoryHandler(cateService)
	bookHandler := handler.NewBookHandler(bookService)

	return &Container{
		AuthHandler:     authHandler,
		CategoryHandler: cateHandler,
		BookHandler:     bookHandler,
		UserRepository:  userRepo,
		DB:              db,
	}, nil
}
