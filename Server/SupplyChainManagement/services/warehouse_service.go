package services

import (
	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/models"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/repositories"
)

func CreateWarehouse(request dto.WarehouseRequest) error {

	warehouse := models.Warehouse{
		Name:    request.Name,
		Address: request.Address,
	}

	return repositories.CreateWarehouse(&warehouse)
}

func GetWarehouses() ([]models.Warehouse, error) {

	return repositories.GetWarehouses()
}

func UpdateWarehouse(id string, request dto.WarehouseRequest) error {

	warehouseID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	warehouse := models.Warehouse{
		Name:    request.Name,
		Address: request.Address,
	}

	warehouse.ID = warehouseID

	return repositories.UpdateWarehouse(&warehouse)
}

func DeleteWarehouse(id string) error {

	warehouseID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return repositories.DeleteWarehouse(warehouseID)
}
