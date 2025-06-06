package service

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"backend/internal/request"
	"errors"

	"github.com/google/uuid"
)

type authServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
	}
}

func (s *authServiceImpl) GetMe() string {
	return "Hế lô anh em"
}

func (s *authServiceImpl) SignUp(req request.ReqSignUp) (*entity.User, error) {
	exists, err := s.userRepo.CheckExistsUserByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("Email đã tồn tại")
	}

	newUser := &entity.User{
		ID:       uuid.NewString(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}
	if err = newUser.HashPassword(); err != nil {
		return nil, errors.New("lỗi băm mật khẩu")
	}

	if err = s.userRepo.CreateUser(newUser); err != nil {
		return nil, errors.New("Đéo tạo được duma")
	}

	return newUser, nil
}
