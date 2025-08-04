package router

import (
	"backend/internal/entity"
	"backend/internal/handler"
	"backend/internal/middleware"
	repository "backend/internal/repository/user"

	"github.com/gin-gonic/gin"
)

func SetupBookRoute(router *gin.RouterGroup, bookHandler *handler.BookHandler, userRepo repository.UserRepository, secretKey string) {
	book := router.Group("/book")
	{
		book.POST("/create", middleware.AuthMiddleware(secretKey, userRepo), middleware.RoleMiddleware(entity.RoleAdmin), bookHandler.CreateBook)
		book.GET("/:id", bookHandler.GetBookByID)
		book.GET("/", bookHandler.GetAllBooks)
		book.PUT("/:id", middleware.AuthMiddleware(secretKey, userRepo), middleware.RoleMiddleware(entity.RoleAdmin), bookHandler.UpdateBook)
		book.DELETE("/:id", middleware.AuthMiddleware(secretKey, userRepo), middleware.RoleMiddleware(entity.RoleAdmin), bookHandler.DeleteBook)
	}
}
