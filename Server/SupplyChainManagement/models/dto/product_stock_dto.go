package dto

type ProductStockRequest struct {
	ProductID   string `json:"product_id" binding:"required"`
	WarehouseID string `json:"warehouse_id" binding:"required"`
	Stock       int `json:"stock"`
}
