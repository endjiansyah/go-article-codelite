package category

import "gorm.io/gorm"

type CategoryRepo interface {
	GetAll() ([]Category, error)
	GetById(ID int) (Category, error)
	Create(category Category) (Category, error)
	Update(category Category) (Category, error)
	Delete(category Category) (Category, error)
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

func (repo *repository) GetById(ID int) (Category, error) {
	var category Category

	err := repo.db.Find(&category, ID).Error
	return category, err
}

func (repo *repository) Create(category Category) (Category, error) {
	err := repo.db.Create(&category).Error
	return category, err
}

func (repo *repository) Update(category Category) (Category, error) {
	err := repo.db.Save(&category).Error
	return category, err
}

func (repo *repository) Delete(category Category) (Category, error) {
	err := repo.db.Delete(&category).Error
	return category, err
}
