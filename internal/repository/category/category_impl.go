package repository

import (
	"backend/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (r *CategoryRepositoryImpl) GetCategoryByID(id uint) (*entity.Category, error) {
	var category entity.Category

	err := r.db.Where("id = ?", id).First(&category).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepositoryImpl) GetAllCategories() ([]entity.Category, error) {
	var categories []entity.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepositoryImpl) CreateCategory(category *entity.Category) (*entity.Category, error) {
	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *CategoryRepositoryImpl) DeleteCategory(id uint) error {
	result := r.db.Where("id = ?", id).Delete(&entity.Category{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}

func (r *CategoryRepositoryImpl) UpdateCategory(category *entity.Category) (*entity.Category, error) {
	if err := r.db.Save(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}
