package repository

import "backend/internal/entity"

type OtpReposiory interface {
	CreateOTP(otp *entity.OTP) error
	VerifyOTP(email string, code string) (bool, error)
	DeleteOTP(email string) error
}
