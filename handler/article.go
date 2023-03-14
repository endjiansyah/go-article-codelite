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
	categorynya := c.Request.URL.Query().Get("category")
	pagenya := c.Request.URL.Query().Get("page")
	limitnya := c.Request.URL.Query().Get("limit")

	Category, err := strconv.Atoi(categorynya)
	if err != nil {
		Category = 0
	}
	Page, err := strconv.Atoi(pagenya)
	if err != nil {
		Page = 0
	}

	Limit, err := strconv.Atoi(limitnya)
	if err != nil {
		Limit = 0
	}

	articles, err := handler.articleService.GetAll(int(Category), int(Page), int(Limit))
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
	CategoryID := c.PostForm("CategoryID")
	categoryID, err := strconv.Atoi(CategoryID)
	if err != nil {
		categoryID = 0
		fmt.Println(err)
		return
	}
	Filename := ""
	Media, err := c.FormFile("Media")
	if err == nil {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(100000)
		Filename = strconv.Itoa(randNum) + filepath.Ext(Media.Filename)

		filename := fmt.Sprintf("uploads/codelite_%s", Filename)
		if err := c.SaveUploadedFile(Media, filename); err != nil {
			c.JSON(500, gin.H{
				"error": "Failed to save media",
			})
			return
		}

	}
	articleRequest := article.ArticleRequest{Title: Title, Media: Filename, Content: Content, CategoryID: int(categoryID)}
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
	}
	if cst.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data tidak ditemukan",
		})
		return
	}
	Title := c.PostForm("Title")
	Content := c.PostForm("Content")
	CategoryID := c.PostForm("CategoryID")
	categoryID, err := strconv.Atoi(CategoryID)
	if err != nil {
		categoryID = 0
	}
	Filename := ""

	Media, errmedia := c.FormFile("Media")
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
			if cst.Media != "" {
				err := os.Remove(cst.Media)
				if err != nil {
					fmt.Println("Error deleting file:", err)
				}
			}
		}

	}

	articleRequest := article.ArticleUpdateRequest{Title: Title, Media: Filename, Content: Content, CategoryID: int(categoryID)}

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
		if cst.Media != "" {
			err := os.Remove(cst.Media)
			if err != nil {
				fmt.Println("Error deleting file:", err)
				return
			}
		}
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
		ID:         cst.ID,
		Title:      cst.Title,
		Content:    cst.Content,
		Media:      cst.Media,
		CategoryID: cst.CategoryID,
		CreatedAt:  cst.CreatedAt,
		UpdatedAt:  cst.UpdatedAt,
	}
}
