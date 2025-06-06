package service

import (
	"backend/internal/entity"
	"backend/internal/request"
)

type AuthService interface {
	GetMe() string
	SignUp(req request.ReqSignUp) (*entity.User, error)
}
