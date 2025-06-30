package repository

import "backend/internal/entity"

type UserRepository interface {
	CheckExistsByEmail(email string) (bool, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(userData *entity.User) error
	UpdateUser(user *entity.User) error
}
