package dto

type StockMovementRequest struct {
	ProductID     string `json:"product_id" binding:"required"`
	WarehouseID   string `json:"warehouse_id" binding:"required"`
	Type          string `json:"type" binding:"required"`
	Qty           int    `json:"qty" binding:"required"`
	ReferenceType string `json:"reference_type"`
	ReferenceID   string `json:"reference_id"`
	Notes         string `json:"notes"`
}
