package repositories

import (
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/models"
	"gorm.io/gorm/clause"
)

func SaveOrIncrementStock(stock *models.ProductStock) error {

	return config.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "product_id"}, {Name: "warehouse_id"}},
		DoUpdates: clause.Assignments(map[string]any{
			"stock": clause.Expr{
				SQL: "product_stock.stock + EXCLUDED.stock",
			},
			"updated_at": clause.Expr{
				SQL: "CURRENT_TIMESTAMP",
			},
		}),
	}).Create(stock).Error
}

func GetProductStocks() ([]models.ProductStock, error) {

	var stocks []models.ProductStock

	err := config.DB.
		Preload("Product").
		Preload("Product.Category").
		Preload("Warehouse").
		Order("created_at DESC").
		Find(&stocks).Error

	return stocks, err
}
