package article

import (
	"gorm.io/gorm"
)

type ArticleRepo interface {
	GetAll(Category int, Page int, Limit int) ([]Article, error)
	GetById(ID int) (Article, error)
	GetMediaId(ID int) (Media, error)
	GetMediaById(ID int) ([]Media, error)
	Create(article Article) (Article, error)
	CreateMedia(media Media) (Media, error)
	Update(article Article) (Article, error)
	Delete(article Article) (Article, error)
	DeleteMedia(media Media) (Media, error)
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
	if Offset <= 0 {
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

func (repo *repository) GetMediaId(ID int) (Media, error) {
	var media Media

	err := repo.db.Find(&media, ID).Error
	return media, err
}

func (repo *repository) GetMediaById(ID int) ([]Media, error) {
	var media []Media

	err := repo.db.Find(&media, "article_id = ?", ID).Error
	return media, err
}

func (repo *repository) Create(article Article) (Article, error) {
	err := repo.db.Create(&article).Error
	return article, err
}

func (repo *repository) CreateMedia(media Media) (Media, error) {
	err := repo.db.Create(&media).Error
	return media, err
}

func (repo *repository) Update(article Article) (Article, error) {
	err := repo.db.Save(&article).Error
	return article, err
}

func (repo *repository) Delete(article Article) (Article, error) {
	err := repo.db.Delete(&article).Error
	return article, err
}

func (repo *repository) DeleteMedia(media Media) (Media, error) {
	err := repo.db.Delete(&media).Error
	return media, err
}
