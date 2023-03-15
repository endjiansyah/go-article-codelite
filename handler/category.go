package handler

import (
	"fmt"
	"go-article-codelite/category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func (handler *categoryHandler) CategoryStore(c *gin.Context) {
	var categoryRequest category.CategoryRequest
	err := c.Bind(&categoryRequest)
	if err != nil {

		listpesaneror := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			pesaneror := fmt.Sprintf("error on %s, because %s", e.Field(), e.ActualTag())
			listpesaneror = append(listpesaneror, pesaneror)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": listpesaneror,
		})
		return
	}
	category, err := handler.categoryService.Create(categoryRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success save data",
		"data":    category,
	})
}

func (handler *categoryHandler) CategoryUpdate(c *gin.Context) {

	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)
	cst, err := handler.categoryService.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	} else if cst.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data not found",
		})
	} else {

		var categoryRequest category.CategoryUpdateRequest

		err := c.Bind(&categoryRequest)
		if err != nil {

			listpesaneror := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				pesaneror := fmt.Sprintf("error on %s, because %s", e.Field(), e.ActualTag())
				listpesaneror = append(listpesaneror, pesaneror)
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"errors": listpesaneror,
			})
			return
		}

		category, err := handler.categoryService.Update(id, categoryRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "success save data",
			"data":    category,
		})
	}

}

func (handler *categoryHandler) CategoryDelete(c *gin.Context) {
	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)

	cst, err := handler.categoryService.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	} else if cst.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data not found",
		})
	} else {
		cst, err := handler.categoryService.Delete(int(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"errors": err,
			})
			return
		}
		categoryResponse := responseCategory(cst)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "success delete category",
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
