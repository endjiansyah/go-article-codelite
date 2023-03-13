package handler

import (
	"go-article-codelite/category"
	"net/http"
	"strconv"

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
func (handler *categoryHandler) CategoryByID(c *gin.Context) {
	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)

	cst, err := handler.categoryService.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
	} else if cst.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data not found",
		})
	} else {
		categoryResponse := responseCategory(cst)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Category with ID : " + c.Param("id"),
			"data":    categoryResponse,
		})
	}
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
