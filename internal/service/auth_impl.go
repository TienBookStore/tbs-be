package service

import (
	"backend/internal/entity"
	customErrors "backend/internal/errors"
	otpRepo "backend/internal/repository/otp"
	userRepo "backend/internal/repository/user"
	"backend/internal/request"
	"backend/internal/utils"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type authServiceImpl struct {
	userRepo userRepo.UserRepository
	otpRepo  otpRepo.OtpReposiory
}

func NewAuthService(userRepo userRepo.UserRepository, otpRepo otpRepo.OtpReposiory) AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
		otpRepo:  otpRepo,
	}
}

func (s *authServiceImpl) GetMe() string {
	return "Hế lô anh em"
}

func (s *authServiceImpl) SignUp(req request.ReqSignUp) (*entity.User, error) {
	exists, err := s.userRepo.CheckExistsByEmail(req.Email)

	if err != nil {
		return nil, fmt.Errorf("Kiểm tra email tồn tại thất bại", err)
	}

	if exists {
		return nil, customErrors.ErrorEmailExists
	}

	newUser := &entity.User{
		ID:       uuid.NewString(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}

	if err = newUser.HashPassword(); err != nil {
		return nil, errors.New("Lỗi băm mật khẩu")
	}

	if err = s.userRepo.CreateUser(newUser); err != nil {
		return nil, errors.New("Không tạo được")
	}

	if err != nil {
		return nil, err
	}

	codeOTP, err := utils.GenerateOTP(6)

	if err != nil {
		return nil, fmt.Errorf("Không tạo được otp", err)
	}

	newOTP := &entity.OTP{
		Email: req.Email,
		Code:  codeOTP,
	}

	if err := s.otpRepo.CreateOTP(newOTP); err != nil {
		return nil, err
	}

	subject := "Mã OTP xác thực tài khoản"
	body := fmt.Sprintf("Mã OTP của bạn là: %s. Mã có hiệu lực 5 phút.", codeOTP)

	if err := utils.SendEmail(req.Email, subject, body); err != nil {
		return nil, err
	}

	return newUser, nil
}
