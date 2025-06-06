package repository

import "backend/internal/entity"

type UserRepository interface {
	//GetUserByEmail(email string) (string, error)
	CheckExistsUserByEmail(email string) (bool, error)
	CreateUser(userData *entity.User) error
}
