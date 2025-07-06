package book

import (
	"backend/internal/entity"
)

type BookRepository interface {
	CreateBook(book *entity.Book) (*entity.Book, error)
	GetAllBooks() ([]entity.Book, error)
	GetBookByID(id uint) (*entity.Book, error)
	UpdateBook(book *entity.Book) (*entity.Book, error)
	DeleteBook(id uint) error
	SearchByTitle(title string) ([]entity.Book, error)
	GetBooksByCategoryID(categoryID uint) ([]entity.Book, error)
}
