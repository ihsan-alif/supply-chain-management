package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/models"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/repositories"
)

func CreateProductStock(request dto.ProductStockRequest) error {

	productUUID, err := uuid.Parse(request.ProductID)
	if err != nil {
		return err
	}

	warehouseUUID, err := uuid.Parse(request.WarehouseID)
	if err != nil {
		return err
	}

	_, err = repositories.GetProductByID(productUUID)
	if err != nil {
		return errors.New("product not found")
	}

	_, err = repositories.GetWarehouseByID(warehouseUUID)
	if err != nil {
		return errors.New("warehouse not found")
	}

	stock := models.ProductStock{
		ProductID:   productUUID,
		WarehouseID: warehouseUUID,
		Stock:       request.Stock,
	}

	return repositories.SaveOrIncrementStock(nil, &stock)
}

func GetProductStocks() ([]models.ProductStock, error) {

	return repositories.GetProductStocks()
}
