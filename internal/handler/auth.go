package handler

import (
	"backend/internal/request"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) GetMe(c *gin.Context) {
	result := h.authService.GetMe()
	c.JSON(http.StatusOK, gin.H{"message": result})
}

func (h *AuthHandler) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req request.ReqSignUp

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "...",
		})
		return
	}
	
	user, err := h.authService.SignUp(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "succesful",
		"user":    user,
	})
}
