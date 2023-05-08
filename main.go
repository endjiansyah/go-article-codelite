package main

import (
	"fmt"
	"go-article-codelite/article"
	"go-article-codelite/category"
	"go-article-codelite/handler"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Mengatur konfigurasi Viper
	viper.SetConfigType("yaml")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file")
	}
	dsn := viper.GetString("DB.HOST")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println(postgres.Open(dsn))
	if err != nil {
		log.Fatal("gagal nyambung DB")
	}

	db.AutoMigrate(&category.Category{})
	db.AutoMigrate(&article.Article{})
	db.AutoMigrate(&article.Media{})

	categoryRepo := category.NewRepository(db)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	articleRepo := article.NewRepository(db)
	articleService := article.NewService(articleRepo)
	articleHandler := handler.NewArticleHandler(articleService)

	router := gin.Default()
	router.GET("/category", categoryHandler.ListCategory)
	router.GET("/category/:id", categoryHandler.CategoryByID)
	router.POST("/category", categoryHandler.CategoryStore)
	router.PUT("/category/:id", categoryHandler.CategoryUpdate)
	router.DELETE("/category/:id", categoryHandler.CategoryDelete)

	router.GET("/article", articleHandler.ListArticle)
	router.GET("/article/:id", articleHandler.ArticleByID)
	router.POST("/article", articleHandler.ArticleStore)
	router.PUT("/article/:id", articleHandler.ArticleUpdate)
	router.DELETE("/article/:id", articleHandler.ArticleDelete)

	router.POST("/article/:id/media", articleHandler.ArticleMediaCreate)
	router.GET("/media/:id", articleHandler.MediaByID)
	router.PUT("/media/:id", articleHandler.MediaUpdate)
	router.DELETE("/media/:id", articleHandler.MediaDelete)

	router.Static("/uploads/", "./uploads")
	router.Run()
}
