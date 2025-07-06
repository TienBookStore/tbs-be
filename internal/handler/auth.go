package handler

import (
	"backend/internal/request"
	service "backend/internal/service/auth"

	"backend/internal/utils"
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
			"messeage": err.Error(),
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
		"message": "Đăng kí thành công, vui lòng kiểm tra email.",
		"user":    user,
	})
}

func (h *AuthHandler) VerifyOTPSignUp(c *gin.Context) {
	var req request.ReqVerifyOTP

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.authService.VerifyOTPSignUp(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "active tài khoản thành công",
		"user":    user,
	})
}

func (h *AuthHandler) ResendOTPSignUp(c *gin.Context) {
	var req request.ReqResendOTP

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.authService.ResendOTP(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Gửi OTP thành công vui lòng kiểm tra email",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.ReqLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.authService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không tạo được token"})
		return
	}

	c.SetCookie(
		"jwt",
		tokenString,
		3600*24,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Đăng nhập thành công",
		"user":    user,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "đăng xuất thành công."})
}
