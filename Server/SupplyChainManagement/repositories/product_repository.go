package repositories

import (
	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/models"
)

func CreateProduct(product *models.Product) error {

	return config.DB.Create(product).Error
}

func GetProducts(limit, offset int, search string) ([]models.Product, int64, error) {

	var products []models.Product
	var total int64

	query := config.DB.Model(&models.Product{})

	if search != "" {
		query = query.Where(
			"(name ILIKE ? OR code ILIKE ?)",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	query.Count(&total)

	err := query.
		Preload("Category").
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&products).Error

	return products, total, err
}

func GetProductByID(id uuid.UUID) (*models.Product, error) {

	var product models.Product

	err := config.DB.
		Preload("Category").
		Where("id = ?", id).
		First(&product).Error

	return &product, err
}

func UpdateProduct(product *models.Product) error {

	return config.DB.Model(product).Updates(models.Product{
		CategoryID:   product.CategoryID,
		Code:         product.Code,
		Name:         product.Name,
		Description:  product.Description,
		Unit:         product.Unit,
		ImageURL:     product.ImageURL,
		MinimumStock: product.MinimumStock,
	}).Error
}

func DeleteProduct(id uuid.UUID) error {

	return config.DB.Delete(
		&models.Product{},
		"id = ?",
		id,
	).Error
}
