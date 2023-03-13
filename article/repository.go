package article

import "gorm.io/gorm"

type ArticleRepo interface {
	GetAll() ([]Article, error)
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

func (repo *repository) GetAll() ([]Article, error) {
	var articles []Article

	err := repo.db.Find(&articles).Error
	return articles, err
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
