package handler

import (
	"go-article-codelite/category"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandler struct {
	categoryService category.Service
}

func NewCategoryHandler(categoryService category.Service) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (handler *categoryHandler) ListCategory(c *gin.Context) {
	categories, err := handler.categoryService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	var categoriesResponse []category.CategoryResponse
	for _, cst := range categories {
		categoryResponse := responseCategory(cst)
		categoriesResponse = append(categoriesResponse, categoryResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "List Category",
		"data":    categoriesResponse,
	})
}

func responseCategory(cst category.Category) category.CategoryResponse {
	return category.CategoryResponse{
		ID:          cst.ID,
		Name:        cst.Name,
		Description: cst.Description,
		CreatedAt:   cst.CreatedAt,
		UpdatedAt:   cst.UpdatedAt,
	}
}
