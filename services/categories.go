package services

import (
	"database/sql"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/md-shadhin-mia/echo-inventory/initialize"
	"github.com/md-shadhin-mia/echo-inventory/models"
)

type CategoryService struct {
	db *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{db: initialize.DB}
}

func (cs *CategoryService) GetAllCategories(limit int, offset int) ([]models.Category, error) {
	if limit == 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	var categories []models.Category
	result := cs.db.Offset(offset).Limit(limit).Find(&categories)
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

func (cs *CategoryService) CreateCategory(category *models.Category) (*models.Category, error) {
	result := cs.db.Create(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (cs *CategoryService) UpdateCategory(id uint, category *models.Category) (*models.Category, error) {
	var existingCategory models.Category
	if err := cs.db.First(&existingCategory, id).Error; err != nil {
		return nil, err // Return error if category not found
	}
	result := cs.db.Model(&existingCategory).Updates(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &existingCategory, nil
}

func (cs *CategoryService) DeleteCategory(id uint) error {
	result := cs.db.Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
