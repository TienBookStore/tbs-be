package entity

import (
	"time"
)

const (
	RoleAdmin    = "admin"
	RoleCustomer = "customer"
)

type User struct {
	ID          string    `json:"id" gorm:"type:varchar(36);primaryKey"`
	FullName    string    `json:"full_name" gorm:"type:varchar(150)"`
	Email       string    `json:"email" gorm:"type:varchar(150);uniqueIndex;not null"`
	Role        string    `json:"role" gorm:"type:role;default:'customer';not null"`
	IsActive    bool      `json:"is_active" gorm:"default:false"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
