package entity

import (
	"time"
)

type OTP struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"index;not null"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
