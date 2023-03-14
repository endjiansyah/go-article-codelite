package article

import (
	"gorm.io/gorm"
)

type ArticleRepo interface {
	GetAll(Category int, Page int, Limit int) ([]Article, error)
	GetById(ID int) (Article, error)
	Create(article Article) (Article, error)
	Update(article Article) (Article, error)
	Delete(article Article) (Article, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) GetAll(Category int, Page int, Limit int) ([]Article, error) {
	var articles []Article

	Offset := (Page - 1) * Limit
	if Offset <= 1 {
		Offset = 0
	}
	if Offset == 0 && Limit <= 0 {
		Limit = -1
	}
	if Category <= 0 {
		err := repo.db.Limit(Limit).Offset(Offset).Find(&articles).Error
		return articles, err
	} else {
		err := repo.db.Limit(Limit).Offset(Offset).Find(&articles, "category_id = ?", Category).Error
		return articles, err
	}
}

func (repo *repository) GetById(ID int) (Article, error) {
	var article Article

	err := repo.db.Find(&article, ID).Error
	return article, err
}

func (repo *repository) Create(article Article) (Article, error) {
	err := repo.db.Create(&article).Error
	return article, err
}

func (repo *repository) Update(article Article) (Article, error) {
	err := repo.db.Save(&article).Error
	return article, err
}

func (repo *repository) Delete(article Article) (Article, error) {
	err := repo.db.Delete(&article).Error
	return article, err
}
