package router

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(router *gin.RouterGroup, authHandler *handler.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.GET("/", authHandler.Welcome)
		auth.GET("/me", authHandler.GetMe)
		auth.POST("/sign-up", authHandler.SignUp)
		auth.POST("/sign-up/verify-otp", authHandler.VerifyOTPSignUp)
		auth.POST("/sign-up/resend-otp", authHandler.ResendOTPSignUp)
	}
}
