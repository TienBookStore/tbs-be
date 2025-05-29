package router

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(router *gin.RouterGroup, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.GET("/me", authHandler.GetMe)
	}
}