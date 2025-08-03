package service

import (
	"backend/internal/entity"
	bookRepo "backend/internal/repository/book"
	cateRepo "backend/internal/repository/category"
	"backend/internal/request"
	"errors"

	"github.com/google/uuid"
)

type bookServiceImpl struct {
	bookRepo bookRepo.BookRepository
	cateRepo cateRepo.CategoryRepository
}

func NewBookService(bookRepo bookRepo.BookRepository, cateRepo cateRepo.CategoryRepository) BookService {
	return &bookServiceImpl{
		bookRepo: bookRepo,
		cateRepo: cateRepo,
	}
}

func (s *bookServiceImpl) CreateBook(req request.ReqCreateBook) (*entity.Book, error) {
	var category []entity.Category

	for _, cateID := range req.CategoryIDs {
		cate, err := s.cateRepo.GetCategoryByID(cateID)	
		
		if err != nil {
			return nil, errors.New("failed to get category: " + err.Error())
		}
		
		if cate == nil {
			return nil, errors.New("category not found with ID: " + cateID)
		}
		
		category = append(category, *cate)
	}

	book := &entity.Book{
		ID:          uuid.NewString(),
		Title:       req.Title,
		Quantity:    req.Quantity,
		Type:        req.Type,
		Description: req.Description,
		Supplier:    req.Supplier,
		Price:       req.Price,
		Language:    req.Language,
		Cover:       req.Cover,
		Year:        req.Year,
		PageNumber:  req.PageNumber,
		Categories:  category,
	}

	book, err := s.bookRepo.CreateBook(book)

	if err != nil {
		return nil, errors.New("failed to create book: " + err.Error())
	}

	if book == nil {
		return nil, errors.New("book creation returned nil")
	}

	return book, nil
}
