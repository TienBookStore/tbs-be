package repository

import (
	"backend/internal/entity"
	"errors"
	"time"

	"gorm.io/gorm"
)

type otpReposioryImpl struct {
	db *gorm.DB
}

func NewOtpRepository(db *gorm.DB) OtpReposiory {
	return &otpReposioryImpl{
		db: db,
	}
}

func (r *otpReposioryImpl) CreateOTP(otp *entity.OTP) error {
	if err := r.db.Debug().Create(otp).Error; err != nil {
		return err
	}

	return nil
}

func (r *otpReposioryImpl) DeleteOTP(email string) error {
	result := r.db.Where("email = ?", email).Delete(&entity.OTP{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("không tìm thấy otp")
	}

	return nil
}

func (r *otpReposioryImpl) VerifyOTP(email string, code string) (bool, error) {
	var otp entity.OTP

	if err := r.db.Where("email = ? AND code = ?", email, code).First(&otp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	if otp.ExpiresAt.Before(time.Now()) {
		return false, errors.New("OTP đã hết hiệu lực")
	}

	return true, nil
}
