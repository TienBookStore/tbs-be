package repository

import (
	"backend/internal/entity"
	cateRepo "backend/internal/repository/category"
	req "backend/internal/request"
	"backend/internal/utils"
	"errors"

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

func (s *categoryServiceImpl) GetCategoryByID(id string) (*entity.Category, error) {
	category, err := s.cateRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryServiceImpl) GetAllCategories() ([]*entity.Category, error) {
	categories, err := s.cateRepo.GetAllCategories()

	if err != nil {
		return nil, err
	}

	categoryPtrs := make([]*entity.Category, len(categories))
	for i := range categories {
		categoryPtrs[i] = &categories[i]
	}

	return categoryPtrs, nil
}

func (s *categoryServiceImpl) UpdateCategory(id string, req req.ReqUpdateCategory) (*entity.Category, error) {
	category, err := s.cateRepo.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("category not found")
	}

	category.Name = req.Name
	category.Slug = utils.GenerateSlug(req.Name)

	updatedCategory, err := s.cateRepo.UpdateCategory(category)
	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (s *categoryServiceImpl) DeleteCategory(id string) error {
	err := s.cateRepo.DeleteCategory(id)
	if err != nil {
		return err
	}

	return nil
}
