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
	"time"

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
		Email:     req.Email,
		Code:      codeOTP,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	if err := s.otpRepo.DeleteOTP(req.Email); err != nil {

	}

	if err := s.otpRepo.CreateOTP(newOTP); err != nil {
		return nil, err
	}

	subject := "Mã OTP xác thực tài khoản"
	body := fmt.Sprintf("Mã OTP của bạn là: %s. Mã có hiệu lực 5 phút.", codeOTP)

	if err := utils.SendOTPByEmail(req.Email, subject, body); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *authServiceImpl) VerifyOTPSignUp(req request.ReqVerifyOTP) error {
	otp, err := s.otpRepo.GetOTPByEmail(req.Email)

	if err != nil {
		return err
	}

	if otp == nil {
		return errors.New("OTP không tồn tại hoặc đã hết hạn, vui lòng gửi lại OTP")
	}

	if otp.Code != req.Code {
		return errors.New("Mã OTP không đúng")
	}

	if time.Now().After(otp.ExpiresAt) {
		return errors.New("Mã OTP đã hết hạn, vui lòng gửi lại OTP")
	}

	user, err := s.userRepo.GetUserByEmail(req.Email)

	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("Email không tồn tại")
	}

	if user.IsActive {
		return errors.New("Email đã được kích hoạt")
	}

	user.IsActive = true

	if err := s.userRepo.UpdateUser(user); err != nil {
		return err
	}

	if err := s.otpRepo.DeleteOTP(otp.Email); err != nil {

	}

	return nil
}
