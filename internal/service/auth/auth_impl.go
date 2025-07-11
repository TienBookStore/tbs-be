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

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("Lỗi băm mật khẩu: %v", err)
	}

	newUser := &entity.User{
		ID:       uuid.NewString(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err = s.userRepo.CreateUser(newUser); err != nil {
		return nil, errors.New("Không tạo được")
	}

	codeOTP, err := utils.GenerateOTP(6)

	if err != nil {
		return nil, fmt.Errorf("Không tạo được otp", err)
	}

	newOTP := &entity.OTP{
		ID:        uuid.NewString(),
		Email:     req.Email,
		Code:      codeOTP,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	exists, err = s.otpRepo.CheckExistsOTPByEmail(req.Email)

	if err != nil {
		return nil, fmt.Errorf("Lỗi kiểm tra OTP", err)
	}

	if exists {
		if err := s.otpRepo.DeleteOTP(req.Email); err != nil {
			return nil, fmt.Errorf("Lỗi xóa OTP", err)
		}
	}

	if err := s.otpRepo.CreateOTP(newOTP); err != nil {
		return nil, fmt.Errorf("Không tạo trong db được otp: %v", err)
	}

	subject := "Mã OTP xác thực tài khoản"
	body := fmt.Sprintf("Mã OTP của bạn là: %s. Mã có hiệu lực 5 phút.", codeOTP)

	if err := utils.SendOTPByEmail(req.Email, subject, body); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *authServiceImpl) VerifyOTPSignUp(req request.ReqVerifyOTP) (*entity.User, error) {
	otp, err := s.otpRepo.GetOTPByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if otp == nil {
		return nil, errors.New("OTP không tồn tại hoặc đã hết hạn, vui lòng gửi lại OTP")
	}

	if otp.Code != req.Code {
		return nil, errors.New("Mã OTP không đúng")
	}

	if time.Now().After(otp.ExpiresAt) {
		return nil, errors.New("Mã OTP đã hết hạn, vui lòng gửi lại OTP")
	}

	user, err := s.userRepo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("Email không tồn tại")
	}

	if user.IsActive {
		return nil, errors.New("Email đã được kích hoạt")
	}

	user.IsActive = true

	if err := s.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}

	exists, err := s.otpRepo.CheckExistsOTPByEmail(req.Email)

	if err != nil {
		return nil, fmt.Errorf("Lỗi kiểm tra OTP", err)
	}

	if exists {
		if err := s.otpRepo.DeleteOTP(req.Email); err != nil {
			return nil, fmt.Errorf("Lỗi xóa OTP", err)
		}
	}

	return user, nil
}

func (s *authServiceImpl) ResendOTP(req request.ReqResendOTP) error {
	codeOTP, err := utils.GenerateOTP(6)

	if err != nil {
		return fmt.Errorf("Không tạo được otp", err)
	}

	newOTP := &entity.OTP{
		Email:     req.Email,
		Code:      codeOTP,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	if err := s.otpRepo.DeleteOTP(req.Email); err != nil {

	}

	if err := s.otpRepo.CreateOTP(newOTP); err != nil {
		return err
	}

	subject := "Mã OTP xác thực tài khoản"
	body := fmt.Sprintf("Mã OTP của bạn là: %s. Mã có hiệu lực 5 phút.", codeOTP)

	if err := utils.SendOTPByEmail(req.Email, subject, body); err != nil {
		return err
	}

	return nil
}

func (s *authServiceImpl) Login(req request.ReqLogin) (*entity.User, error) {
	user, err := s.userRepo.GetUserByEmail(req.Email)

	if err != nil {
		return nil, fmt.Errorf("Lỗi khi lấy thông tin người dùng: %v", err)
	}

	if user == nil {
		return nil, customErrors.ErrorUserNotFound
	}

	if !user.IsActive {
		return nil, customErrors.ErrorUserNotActive
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, customErrors.ErrorInvalidCredentials
	}

	return user, nil
}

func (s *authServiceImpl) ChangePassword(user *entity.User, req request.ReqChangePassword) error {
	if !utils.CheckPassword(req.OldPassword, user.Password) {
		return customErrors.ErrorOldPasswordIncorrect
	}

	if req.OldPassword == req.NewPassword {
		return errors.New("Mật khẩu mới không được trùng với mật khẩu cũ")
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return fmt.Errorf("Lỗi băm mật khẩu: %v", err)
	}
	user.Password = hashedPassword
	if err := s.userRepo.UpdateUser(user); err != nil {
		return customErrors.ErrorUserUpdateFailed
	}

	return nil
}
