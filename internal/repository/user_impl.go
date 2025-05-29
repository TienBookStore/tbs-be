package repository

import "gorm.io/gorm"

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (string, error) {
	return "Làm gì đó đi", nil
}