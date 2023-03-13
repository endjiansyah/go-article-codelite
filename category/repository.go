package category

import "gorm.io/gorm"

type CategoryRepo interface {
	GetAll() ([]Category, error)
	GetById(ID int) (Category, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Category, error) {
	var categories []Category

	err := repo.db.Find(&categories).Error
	return categories, err
}

func (repo *repository) GetById(ID int) (Category, error) { //getById
	var category Category

	err := repo.db.Find(&category, ID).Error
	return category, err
}
