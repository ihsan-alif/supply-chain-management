package repositories

import (
	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/models"
)

func CreateWarehouse(warehouse *models.Warehouse) error {

	return config.DB.Create(warehouse).Error
}

func GetWarehouses() ([]models.Warehouse, error) {

	var warehouses []models.Warehouse

	err := config.DB.
		Order("created_at DESC").
		Find(&warehouses).Error

	return warehouses, err
}

func GetWarehouseByID(id uuid.UUID) (*models.Warehouse, error) {

	var warehouse models.Warehouse

	err := config.DB.
		Where("id = ?", id).
		First(&warehouse).Error

	return &warehouse, err
}

func UpdateWarehouse(warehouse *models.Warehouse) error {

	return config.DB.Model(warehouse).Updates(models.Warehouse{
		Name:    warehouse.Name,
		Address: warehouse.Address,
	}).Error
}

func DeleteWarehouse(id uuid.UUID) error {

	return config.DB.Delete(
		&models.Warehouse{},
		"id = ?",
		id,
	).Error
}
