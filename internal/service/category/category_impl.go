package repository

import (
	"backend/internal/entity"
	cateRepo "backend/internal/repository/category"
	"backend/internal/utils"

	"github.com/google/uuid"
)

type categoryServiceImpl struct {
	cateRepo cateRepo.CategoryRepository
}

func NewCategoryService(cateRepo cateRepo.CategoryRepository) CategoryService {
	return &categoryServiceImpl{
		cateRepo: cateRepo,
	}
}

func (s *categoryServiceImpl) CreateCategory(name string) (*entity.Category, error) {
	category := entity.Category{
		ID:   uuid.NewString(),
		Name: name,
		Slug: utils.GenerateSlug(name),
	}

	createdCategory, err := s.cateRepo.CreateCategory(&category)

	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}
