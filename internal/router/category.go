package router

import (
	"backend/internal/entity"
	"backend/internal/handler"
	"backend/internal/middleware"
	repository "backend/internal/repository/user"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoute(router *gin.RouterGroup, cateHandler *handler.CategoryHandler, userRepo repository.UserRepository, secretKey string) {
	category := router.Group("/categories")
	{
		category.POST("/create", middleware.AuthMiddleware(secretKey, userRepo), middleware.RoleMiddleware(entity.RoleAdmin), cateHandler.Create)
	}
}
