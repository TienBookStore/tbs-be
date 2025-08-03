package repository

import (
	"backend/internal/entity"
)

type BookRepository interface {
	CreateBook(book *entity.Book) (*entity.Book, error)
	GetAllBooks() ([]entity.Book, error)
	GetBookByID(id string) (*entity.Book, error)
	UpdateBook(book *entity.Book) (*entity.Book, error)
	DeleteBook(id string) error
	SearchByTitle(title string) ([]entity.Book, error)
	GetBooksByCategoryID(categoryID string) ([]entity.Book, error)
}
