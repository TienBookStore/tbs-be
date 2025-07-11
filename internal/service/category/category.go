package repository

import (
	"backend/internal/entity"
	"backend/internal/request"
)

type CategoryService interface {
	CreateCategory(name string) (*entity.Category, error)
	GetCategoryByID(id string) (*entity.Category, error)
	GetAllCategories() ([]*entity.Category, error)
	UpdateCategory(id string, req request.ReqUpdateCategory) (*entity.Category, error)
	DeleteCategory(id string) error
}
