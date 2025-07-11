package repository

import (
	"backend/internal/entity"
)

type CategoryRepository interface {
	GetCategoryByID(id string) (*entity.Category, error)
	GetAllCategories() ([]entity.Category, error)
	CreateCategory(category *entity.Category) (*entity.Category, error)
	DeleteCategory(id string) error
	UpdateCategory(category *entity.Category) (*entity.Category, error)
}
