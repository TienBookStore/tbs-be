package repository

import "backend/internal/entity"

type UserRepository interface {
	//GetUserByEmail(email string) (string, error)
	CheckExistsByEmail(email string) (bool, error)
	GetUserByEmail(email string) (*entity.User, error)
	CreateUser(userData *entity.User) error
	ActivateUser(email string) error
}
