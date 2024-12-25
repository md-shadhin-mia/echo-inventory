package services

import (
	"database/sql"
	"errors"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{db: db}
}

func (cs *CategoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := cs.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (cs *CategoryService) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	result := cs.db.First(&category, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, sql.ErrNoRows
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (cs *CategoryService) CreateCategory(category *models.Category) error {
	result := cs.db.Create(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cs *CategoryService) UpdateCategory(category *models.Category) error {
	result := cs.db.Save(category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cs *CategoryService) DeleteCategory(id uint) error {
	result := cs.db.Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}