package entity

import (
	"time"
)

type OTP struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primaryKey"`
	Email     string    `gorm:"index;not null"`
	Code      string    `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}
