package main

import (
	"go-article-codelite/article"
	"go-article-codelite/category"
	"log"

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
	db.AutoMigrate(&article.Article{})

}
