package repository

import (
	"backend/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) CheckExistsUserByEmail(email string) (bool, error) {
	user := entity.User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *userRepositoryImpl) CreateUser(userData *entity.User) error {
	return r.db.Debug().Create(userData).Error
}
