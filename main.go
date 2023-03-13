package main

import (
	"go-article-codelite/category"
	"go-article-codelite/handler"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=goarticle-codelite port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("gagal nyambung DB")
	}

	db.AutoMigrate(&category.Category{})
	// db.AutoMigrate(&article.Article{})

	categoryRepo := category.NewRepository(db) // panggil repository category (mirip model di CI)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router := gin.Default()
	router.GET("/category", categoryHandler.ListCategory)
	router.GET("/category/:id", categoryHandler.CategoryByID)
	router.Run()
}
