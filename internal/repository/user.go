package repository

type UserRepository interface {
	GetUserByEmail(email string) (string, error)
}