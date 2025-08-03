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

func (s *bookServiceImpl) GetBookByID(id string) (*entity.Book, error) {
	book, err := s.bookRepo.GetBookByID(id)

	if err != nil {
		return nil, errors.New("failed to get book by ID: " + err.Error())
	}

	if book == nil {
		return nil, nil // Book not found
	}

	return book, nil
}

func (s *bookServiceImpl) GetAllBooks() ([]entity.Book, error) {
	books, err := s.bookRepo.GetAllBooks()

	if err != nil {
		return nil, errors.New("failed to get all books: " + err.Error())
	}

	if books == nil {
		return nil, nil // No books found
	}

	return books, nil
}
