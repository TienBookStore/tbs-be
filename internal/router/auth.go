package router

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	repository "backend/internal/repository/user"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(router *gin.RouterGroup, authHandler *handler.AuthHandler, userRepo repository.UserRepository, secretKey string) {
	auth := router.Group("/auth")
	{
		auth.GET("/", authHandler.Welcome)
		auth.GET("/me", middleware.AuthMiddleware(secretKey, userRepo), authHandler.GetMe)
		auth.POST("/sign-up", authHandler.SignUp)
		auth.POST("/sign-up/verify-otp", authHandler.VerifyOTPSignUp)
		auth.POST("/login", authHandler.Login)
		auth.POST("/sign-up/resend-otp", authHandler.ResendOTPSignUp)
		auth.POST("/logout", authHandler.Logout)
		auth.PUT("/change-password", middleware.AuthMiddleware(secretKey, userRepo), authHandler.ChangePassword)
	}
}
