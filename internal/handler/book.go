package handler

import (
	"backend/internal/request"
	service "backend/internal/service/book"
	"backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{
		bookService: bookService,
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var req request.ReqCreateBook

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	book, err := h.bookService.CreateBook(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status: http.StatusOK,
		Message: "Create book successfully",
		Data: book,
	})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	book, err := h.bookService.GetBookByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	if book == nil {
		c.JSON(http.StatusNotFound, utils.Response{
			Status: http.StatusNotFound,
			Message: "Book not found",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status: http.StatusOK,
		Message: "Get book successfully",
		Data: book,
	})
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status: http.StatusOK,
		Message: "Get all books successfully",
		Data: books,
	})
}
 
func (h *BookHandler) UpdateBook(c *gin.Context) {
	var req request.ReqUpdateBook
	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{
			Status: http.StatusBadRequest,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	book, err := h.bookService.UpdateBook(id, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.Response{
		Status: http.StatusOK,
		Message: "Update book successfully",
		Data: book,
	})
}