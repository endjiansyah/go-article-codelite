package handler

import (
	"fmt"
	"go-article-codelite/article"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type articleHandler struct {
	articleService article.Service
}

func NewArticleHandler(articleService article.Service) *articleHandler {
	return &articleHandler{articleService}
}

func (handler *articleHandler) ListArticle(c *gin.Context) {
	articles, err := handler.articleService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	var articlesResponse []article.ArticleResponse
	for _, cst := range articles {
		articleResponse := responseArticle(cst)
		articlesResponse = append(articlesResponse, articleResponse)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "List Article",
		"data":    articlesResponse,
	})
}
func (handler *articleHandler) ArticleByID(c *gin.Context) {
	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)

	cst, err := handler.articleService.GetById(int(id))

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
		articleResponse := responseArticle(cst)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Article with ID : " + c.Param("id"),
			"data":    articleResponse,
		})
	}
}

func (handler *articleHandler) ArticleStore(c *gin.Context) {
	var articleRequest article.ArticleRequest
	err := c.Bind(&articleRequest)
	if err != nil {

		listpesaneror := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			pesaneror := fmt.Sprintf("error di %s, karena %s", e.Field(), e.ActualTag())
			listpesaneror = append(listpesaneror, pesaneror)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": listpesaneror,
		})
		return
	}
	article, err := handler.articleService.Create(articleRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data tersimpan",
		"data":    article,
	})
}

func (handler *articleHandler) ArticleUpdate(c *gin.Context) {

	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)
	cst, err := handler.articleService.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	} else if cst.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data tidak ditemukan",
		})
	} else {

		var articleRequest article.ArticleUpdateRequest

		err := c.Bind(&articleRequest)
		if err != nil {

			listpesaneror := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				pesaneror := fmt.Sprintf("error di %s, karena %s", e.Field(), e.ActualTag())
				listpesaneror = append(listpesaneror, pesaneror)
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"errors": listpesaneror,
			})
			return
		}

		article, err := handler.articleService.Update(id, articleRequest)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Data tersimpan",
			"data":    article,
		})
	}

}

func (handler *articleHandler) ArticleDelete(c *gin.Context) {
	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)

	cst, err := handler.articleService.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	} else if cst.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data tidak ditemukan",
		})
	} else {
		cst, err := handler.articleService.Delete(int(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": false,
				"errors": err,
			})
			return
		}
		articleResponse := responseArticle(cst)
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "Hapus Article",
			"data":    articleResponse,
		})
	}
}

func responseArticle(cst article.Article) article.ArticleResponse {
	return article.ArticleResponse{
		ID:        cst.ID,
		Title:     cst.Title,
		Content:   cst.Content,
		Media:     cst.Media,
		CreatedAt: cst.CreatedAt,
		UpdatedAt: cst.UpdatedAt,
	}
}
