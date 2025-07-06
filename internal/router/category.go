package router

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoute(router *gin.RouterGroup, cateHandler *handler.CategoryHandler) {
	category := router.Group("/categories")
	{
		category.POST("/create", cateHandler.Create)
	}
}
