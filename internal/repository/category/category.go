package repository

import (
	"backend/internal/entity"
)

type CategoryRepository interface {
	GetCategoryByID(id uint) (*entity.Category, error)
	GetAllCategories() ([]entity.Category, error)
	CreateCategory(category *entity.Category) (*entity.Category, error)
	DeleteCategory(id uint) error
	UpdateCategory(category *entity.Category) (*entity.Category, error)
}
