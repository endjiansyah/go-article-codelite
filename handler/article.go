package handler

import (
	"fmt"
	"go-article-codelite/article"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

	Title := c.PostForm("Title")
	Content := c.PostForm("Content")
	Media, err := c.FormFile("Media")

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to upload file",
		})
		return
	}

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(100000)
	fileName := strconv.Itoa(randNum) + filepath.Ext(Media.Filename)

	filename := fmt.Sprintf("uploads/codelite_%s", fileName)
	if err := c.SaveUploadedFile(Media, filename); err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to save media",
		})
		return
	}

	articleRequest := article.ArticleRequest{Title: Title, Media: filename, Content: Content}

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
		Title := c.PostForm("Title")
		Content := c.PostForm("Content")
		Media, errmedia := c.FormFile("Media")
		Filename := ""

		if errmedia == nil {
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(100000)
			fileName := strconv.Itoa(randNum) + filepath.Ext(Media.Filename)

			Filename = fmt.Sprintf("uploads/codelite_%s", fileName)
			if err := c.SaveUploadedFile(Media, Filename); err != nil {
				c.JSON(500, gin.H{
					"error": "Failed to save media",
				})
				return
			} else {
				err := os.Remove(cst.Media)
				if err != nil {
					fmt.Println("Error deleting file:", err)
					return
				}
			}

		}

		articleRequest := article.ArticleUpdateRequest{Title: Title, Media: Filename, Content: Content}

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
