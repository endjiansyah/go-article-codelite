package handler

import (
	"fmt"
	"go-article-codelite/article"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

	art, err := handler.articleService.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	}
	if art.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data not found",
		})
		return
	}

	articleResponse := responseArticle(art)

	med, _ := handler.articleService.GetMediaById(int(id))
	var mediaResponses []article.MediaResponse
	for _, cst := range med {
		mediaresponse := responseMedia(cst)
		mediaResponses = append(mediaResponses, mediaresponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Article with ID : " + c.Param("id"),
		"data":    articleResponse,
		"Media":   mediaResponses,
	})

}

func (handler *articleHandler) ArticleMediaCreate(c *gin.Context) {
	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)

	art, err := handler.articleService.GetById(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	}
	if art.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "article not found",
		})
		return
	}

	var Filename string
	var Mimetype string
	Media, err := c.FormFile("Media")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input Image, Audio, or Video",
		})
		return
	}
	mimetype := Media.Header.Get("Content-Type")
	mime := strings.Split(mimetype, "/")
	if mime[0] == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Input Image, Audio, or Video",
		})
		return
	}
	if mime[0] != "image" && mime[0] != "video" && mime[0] != "audio" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "your uploaded file is " + mime[0] + ", the allowed file is audio, video,& image",
		})
		return
	}

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(100000)
	randomname := strconv.Itoa(randNum) + filepath.Ext(Media.Filename)
	filename := fmt.Sprintf("uploads/%s/codelite_%s", mime[0], randomname)
	if err := c.SaveUploadedFile(Media, filename); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to save media",
		})
		return
	}
	Filename = fmt.Sprintf("%s/%s", c.Request.Host, filename)
	Mimetype = mime[0]

	mediaRequest := article.MediapostRequest{Media: Filename, Type: Mimetype, ArticleID: int(id)}
	media, err := handler.articleService.CreateMedia(mediaRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "success save media",
		"data":    media,
	})

}

func (handler *articleHandler) MediaByID(c *gin.Context) {
	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)

	med, err := handler.articleService.GetMediaId(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	}
	if med.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data not found",
		})
		return
	}

	mediaResponse := responseMedia(med)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Media with ID : " + c.Param("id"),
		"data":    mediaResponse,
	})

}

func (handler *articleHandler) MediaUpdate(c *gin.Context) {
	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)

	med, err := handler.articleService.GetMediaId(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	}
	if med.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data not found",
		})
		return
	}

	var Filename string
	var Mimetype string
	Media, err := c.FormFile("Media")
	if err == nil {
		mimetype := Media.Header.Get("Content-Type")
		mime := strings.Split(mimetype, "/")

		if mime[0] != "image" && mime[0] != "video" && mime[0] != "audio" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "your uploaded file is " + mime[0] + ", the allowed file is audio, video,& image",
			})
			return
		}

		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(100000)
		randomname := strconv.Itoa(randNum) + filepath.Ext(Media.Filename)
		filename := fmt.Sprintf("uploads/%s/codelite_%s", mime[0], randomname)
		if err := c.SaveUploadedFile(Media, filename); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to save media",
			})
			return
		}
		pathimg := strings.Replace(med.Media, c.Request.Host, ".", -1)
		err := os.Remove(pathimg)
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}
		Filename = fmt.Sprintf("%s/%s", c.Request.Host, filename)
		Mimetype = mime[0]
	}
	mediaRequest := article.MediapostRequest{Media: Filename, Type: Mimetype, ArticleID: int(med.ArticleID)}
	media, err := handler.articleService.UpdateMedia(id, mediaRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Update media with id : " + idnya,
		"data":    media,
	})

}

func (handler *articleHandler) MediaDelete(c *gin.Context) {
	idnya := c.Param("id")
	id, _ := strconv.Atoi(idnya)

	med, err := handler.articleService.GetMediaId(int(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	}
	if med.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data not found",
		})
		return
	}

	pathimg := strings.Replace(med.Media, c.Request.Host, ".", -1)
	errdel := os.Remove(pathimg)
	if errdel != nil {
		fmt.Println("Error deleting file:", err)
	}

	delmedia, err := handler.articleService.DeleteMedia(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Delete Media with id : " + idnya,
		"data":    delmedia,
	})

}

func (handler *articleHandler) ArticleStore(c *gin.Context) {

	Title := c.PostForm("Title")
	if Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'Title' field is required",
		})
		return
	}

	Content := c.PostForm("Content")
	if Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'Content' field is required",
		})
		return
	}

	Author := c.PostForm("Author")
	if Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'Author' field is required",
		})
		return
	}

	CategoryID := c.PostForm("CategoryID")
	if CategoryID == "" {
		CategoryID = "0"
	}

	categoryID, err := strconv.Atoi(CategoryID)
	if err != nil {
		categoryID = 0
	}

	if categoryID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "'CategoryID' field is required & must be integer",
		})
		return
	}

	articleRequest := article.ArticleRequest{Title: Title, Content: Content, Author: Author, CategoryID: int(categoryID)}
	articlecreate, err := handler.articleService.Create(articleRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	// --------[media]---------
	Erruploadcount := 0
	Succuploadcount := 0
	form, _ := c.MultipartForm()
	Media := form.File["Media[]"]
	var mediaResponses []article.MediaResponse

	if err == nil {
		for _, file := range Media {
			mimetype := file.Header.Get("Content-Type")
			mime := strings.Split(mimetype, "/")

			if mime[0] != "image" && mime[0] != "video" && mime[0] != "audio" {
				Erruploadcount++
				continue
			}

			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(100000)
			randomname := strconv.Itoa(randNum) + filepath.Ext(file.Filename)

			filename := fmt.Sprintf("uploads/%s/codelite_%s", mime[0], randomname)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				Erruploadcount++
				continue
			}
			fileName := fmt.Sprintf("%s/%s", c.Request.Host, filename)

			mediaRequest := article.MediapostRequest{Media: fileName, Type: mime[0], ArticleID: articlecreate.ID}
			mediareq, err := handler.articleService.CreateMedia(mediaRequest)
			if err != nil {
				Erruploadcount++
				continue
			}
			mediaresponse := responseMedia(mediareq)
			mediaResponses = append(mediaResponses, mediaresponse)
			Succuploadcount++
		}
	}
	// ---------endmedia-------

	articleResponse := responseArticle(articlecreate)

	c.JSON(http.StatusOK, gin.H{
		"status":               true,
		"message":              "Data tersimpan",
		"data":                 articleResponse,
		"media":                mediaResponses,
		"media_upload_success": Succuploadcount,
		"media_upload_fails":   Erruploadcount,
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
			"message": "data not found",
		})
		return
	}
	Title := c.PostForm("Title")
	Content := c.PostForm("Content")
	Author := c.PostForm("Author")
	CategoryID := c.PostForm("CategoryID")
	categoryID, err := strconv.Atoi(CategoryID)
	if err != nil {
		categoryID = 0
	}
	articleRequest := article.ArticleUpdateRequest{Title: Title, Content: Content, Author: Author, CategoryID: int(categoryID)}

	article, err := handler.articleService.Update(id, articleRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Delete Media with id : ",
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
	}
	if cst.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "data not fond",
		})
		return
	}
	med, _ := handler.articleService.GetMediaById(int(id))
	var mediaResponses []article.MediaResponse
	for _, mdia := range med {
		mediaresponse := responseMedia(mdia)
		pathimg := strings.Replace(mdia.Media, c.Request.Host, ".", -1)
		os.Remove(pathimg)
		handler.articleService.DeleteMedia(int(mediaresponse.ID))
		mediaResponses = append(mediaResponses, mediaresponse)
	}

	artdel, err := handler.articleService.Delete(int(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"errors": err,
		})
		return
	}
	articleResponse := responseArticle(artdel)
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Delete Article & Media with id : " + idnya,
		"data":    articleResponse,
		"media":   mediaResponses,
	})

}

func responseArticle(cst article.Article) article.ArticleResponse {
	return article.ArticleResponse{
		ID:         cst.ID,
		Title:      cst.Title,
		Content:    cst.Content,
		Author:     cst.Author,
		CategoryID: cst.CategoryID,
		CreatedAt:  cst.CreatedAt,
		UpdatedAt:  cst.UpdatedAt,
	}
}
func responseMedia(cst article.Media) article.MediaResponse {
	return article.MediaResponse{
		ID:        cst.ID,
		Media:     cst.Media,
		Type:      cst.Type,
		ArticleID: cst.ArticleID,
	}
}
