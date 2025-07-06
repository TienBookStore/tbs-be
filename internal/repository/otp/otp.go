package repository

import (
	"backend/internal/entity"
)

type OtpReposiory interface {
	CheckExistsOTPByEmail(email string) (bool, error)
	GetOTPByEmail(email string) (*entity.OTP, error)
	CreateOTP(otp *entity.OTP) error
	DeleteOTP(email string) error
}
