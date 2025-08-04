package repository

import (
	"backend/internal/entity"
	"errors"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db,
	}
}

func (r *BookRepositoryImpl) GetBookByID(id string) (*entity.Book, error) {
	var book entity.Book

	err := r.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}

func (r *BookRepositoryImpl) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book

	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepositoryImpl) CreateBook(book *entity.Book) (*entity.Book, error) {
	if err := r.db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepositoryImpl) DeleteBook(id string) error {
    var book entity.Book

    // Lấy book ra trước
    if err := r.db.Preload("Categories").First(&book, "id = ?", id).Error; err != nil {
        return err
    }

    // Xóa liên kết many-to-many
    if err := r.db.Model(&book).Association("Categories").Clear(); err != nil {
        return err
    }

    // Xóa book theo ID để tránh lỗi ORDER BY
    if err := r.db.Delete(&entity.Book{}, "id = ?", id).Error; err != nil {
        return err
    }

    return nil
}


func (r *BookRepositoryImpl) UpdateBook(book *entity.Book) (*entity.Book, error) {
	if err := r.db.Save(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepositoryImpl) SearchByTitle(title string) ([]entity.Book, error) {
	var books []entity.Book

	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&books).Error
	if err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, errors.New("no books found with the given title")
	}

	return books, nil
}

func (r *BookRepositoryImpl) GetBooksByCategoryID(categoryID string) ([]entity.Book, error) {
	var books []entity.Book

	err := r.db.Table("books").
		Joins("JOIN category_book ON category_book.book_id = books.id").
		Where("category_book.category_id = ?", categoryID).
		Find(&books).Error
	if err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return nil, errors.New("no books found for the given category")
	}

	return books, nil
}
