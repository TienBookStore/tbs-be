package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id" gorm:"type:varchar(36);primaryKey"`
	FullName  string    `json:"full_name" gorm:"type:varchar(150)"`
	Email     string    `json:"email" gorm:"type:varchar(150);uniqueIndex;not null"`
	IsActive  bool      `json:"is_active" gorm:"default:false"`
	Password  string    `json:"-" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
