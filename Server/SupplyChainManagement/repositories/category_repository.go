package repositories

import (
	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/models"
)

func CreateCategory(category *models.Category) error {

	return config.DB.Create(category).Error
}

func GetCategories(limit, offset int, search string) ([]models.Category, int64, error) {

	var categories []models.Category
	var total int64

	query := config.DB.Model(&models.Category{})

	if search != "" {
		query = query.Where(
			"name ILIKE ?",
			"%"+search+"%",
		)
	}

	query.Count(&total)

	err := query.
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&categories).Error

	return categories, total, err
}

func GetCategoryByID(id uuid.UUID) (*models.Category, error) {

	var category models.Category

	err := config.DB.
		Where("id = ?", id).
		First(&category).Error

	return &category, err
}

func UpdateCategoryName(id uuid.UUID, name string) error {
	
	return config.DB.Model(&models.Category{}).
		Where("id = ?", id).
		Update("name", name).Error
}

func DeleteCategory(id uuid.UUID) error {

	return config.DB.Delete(
		&models.Category{},
		"id = ?",
		id,
	).Error
}