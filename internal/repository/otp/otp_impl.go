package repository

import (
	"backend/internal/entity"
	"errors"

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

func (r *otpReposioryImpl) GetOTPByEmail(email string) (*entity.OTP, error) {
	var otp entity.OTP

	err := r.db.Where("email = ?", email).First(&otp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &otp, nil
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

