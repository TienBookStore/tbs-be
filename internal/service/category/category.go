package repository

import "backend/internal/entity"

type CategoryService interface {
	CreateCategory(name string) (*entity.Category, error)
}
