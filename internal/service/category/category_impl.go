package repository

import (
	"backend/internal/entity"
	cateRepo "backend/internal/repository/category"
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
	var category entity.Category

	category.Name = name

	createdCategory, err := s.cateRepo.CreateCategory(&category)

	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}
