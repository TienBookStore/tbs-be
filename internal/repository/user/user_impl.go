package repository

import (
	"backend/internal/entity"

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

func (r *userRepositoryImpl) CheckExistsByEmail(email string) (bool, error) {
	var count int64

	if err := r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) CreateUser(userData *entity.User) error {
	return r.db.Debug().Create(userData).Error
}

func (r *userRepositoryImpl) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepositoryImpl) GetUserByID(id string) (*entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}
	return &user, nil
}
