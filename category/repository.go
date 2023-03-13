package category

import "gorm.io/gorm"

type CategoryRepo interface {
	GetAll() ([]Category, error)
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
