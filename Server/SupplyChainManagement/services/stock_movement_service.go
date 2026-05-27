package services

import (
	"errors"

	"github.com/google/uuid"
	"github.com/ihsan-alif/supply-chain-management/config"
	"github.com/ihsan-alif/supply-chain-management/models"
	"github.com/ihsan-alif/supply-chain-management/models/dto"
	"github.com/ihsan-alif/supply-chain-management/repositories"
	"gorm.io/gorm"
)

func CreateStockMovement(request dto.StockMovementRequest, userID string) error {

	productUUID, err := uuid.Parse(request.ProductID)
	if err != nil {
		return err
	}

	warehouseUUID, err := uuid.Parse(request.WarehouseID)
	if err != nil {
		return err
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	tx := config.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		}
	}()

	stock, err := repositories.GetProductStock(tx, productUUID, warehouseUUID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if request.Type == "OUT" {
				return errors.New("insufficient stock: data stock does not exist")
			}

			stock = &models.ProductStock{

				ProductID:   productUUID,
				WarehouseID: warehouseUUID,
				Stock:       0,
			}

			stock.ID = uuid.New()
		} else {
			return err
		}
	}

	switch request.Type {
	case "IN":
		stock.Stock += request.Qty
	case "OUT":
		if stock.Stock < request.Qty {
			return errors.New("insufficient stock")
		}
		stock.Stock -= request.Qty
	case "ADJUSTMENT":
		stock.Stock = request.Qty
	default:
		return errors.New("invalid stock movement type")
	}

	err = repositories.UpdateProductStock(tx, stock)
	if err != nil {
		return err
	}

	movement := models.StockMovement{
		ProductID:     productUUID,
		WarehouseID:   warehouseUUID,
		Type:          request.Type,
		Qty:           request.Qty,
		ReferenceType: request.ReferenceType,
		Notes:         request.Notes,
		CreatedBy:     userUUID,
	}

	err = repositories.CreateStockMovement(tx, &movement)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func GetStockMovements() ([]models.StockMovement, error) {

	return repositories.GetStockMovements()
}
