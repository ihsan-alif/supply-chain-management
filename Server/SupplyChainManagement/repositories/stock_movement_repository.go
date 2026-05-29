package repositories

import (
	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/models"
	"gorm.io/gorm"
)

func CreateStockMovement(tx *gorm.DB, movement *models.StockMovement) error {

	return tx.Create(movement).Error
}

func GetProductStock(tx *gorm.DB, productID uuid.UUID, warehouseID uuid.UUID) (*models.ProductStock, error) {

	var stock models.ProductStock

	err := tx.Where(
		"product_id = ? AND warehouse_id = ?",
		productID,
		warehouseID,
	).First(&stock).Error

	return &stock, err
}

func GetStockMovements() ([]models.StockMovement, error) {

	var movements []models.StockMovement

	err := config.DB.
		Preload("Product").
		Preload("Warehouse").
		Preload("User").
		Order("created_at DESC").
		Find(&movements).Error

	return movements, err
}
