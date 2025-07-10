package handler

import (
	"backend/internal/request"
	service "backend/internal/service/category"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req request.ReqCreateCategory

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	category, err := h.categoryService.CreateCategory(req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Category created successfully",
		"category": category,
	})
}

func (h *CategoryHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	category, err := h.categoryService.GetCategoryByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if categories == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No categories found",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})
}

func (h *CategoryHandler) Update(c *gin.Context) {
	var req request.ReqUpdateCategory
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	category, err := h.categoryService.UpdateCategory(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "Category updated successfully",
		"updated_category": category,
	})
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.categoryService.DeleteCategory(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}
