package service

import (
	"backend/internal/entity"
	"backend/internal/request"
)

type BookService interface {
	CreateBook(req request.ReqCreateBook) (*entity.Book, error)
	GetBookByID(id string) (*entity.Book, error)
}
